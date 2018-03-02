/*
Copyright (C) 2017 Verizon. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package interpreter

import (
	"github.com/lavaorg/lrtx/mlog"
	"github.com/lavaorg/lua"
	"github.com/lavaorg/northstar/rte/config"
	"github.com/lavaorg/northstar/rte/rlimit"
	"github.com/lavaorg/northstar/rte/rtepub"
	"strings"
	"time"
)

type LuaInterpreter struct {
	State  *State
	rLimit rlimit.ResourceLimit
}

func NewLuaInterpreter(rLimit rlimit.ResourceLimit) rtepub.Interpreter {
	return &LuaInterpreter{rLimit: rLimit}
}

func (i *LuaInterpreter) DoREPL(input *rtepub.Input) *rtepub.Output {
	timer := Lua.NewTimer("DoREPLTimer")
	startedOn := time.Now()
	mlog.Debug("Running main: %s, code: %s, args: %s, timeout: %d, memory: %v, accountId: %s, "+
		"invocationId: %s", input.MainFn, input.Code, input.Args, input.Timeout, input.Memory,
		input.AccountId, input.InvocationId)

	state, err := CreateState(input)
	if err != nil {
		mlog.Error("Failed to create state: %v", err.Error())
		timer.Stop()
		ErrDoREPL.Incr()
		return &rtepub.Output{StartedOn: startedOn,
			Status:     rtepub.STATE_CREATE_FAILED,
			ErrorDescr: err.Error()}
	}

	i.State = state
	defer state.Close()

	var rLimitErr error
	if config.EnableRLimit {
		msg, err := i.rLimit.StartMonitoring(&rlimit.Resources{Memory: input.Memory})
		if err != nil {
			mlog.Error("Failed to start resource monitoring: %v", err.Error())
			timer.Stop()
			ErrDoREPL.Incr()
			return &rtepub.Output{StartedOn: startedOn,
				Status:     rtepub.START_MONITORING_FAILED,
				ErrorDescr: err.Error()}
		}

		go func() {
		L:
			for {
				select {
				case rLimitErr = <-msg:
					mlog.Debug("Received resource limit msg: %v", rLimitErr)
					i.State.Close()
					break L
				default:
				}
			}

			mlog.Debug("Lua rlimit goroutine stopped")
		}()
	}

	err = state.LuaState.DoString(input.Code)
	if err != nil {
		mlog.Error("DoString error: %v", err)
		i.rLimit.StopMonitoring()
		timer.Stop()
		ErrDoREPL.Incr()
		execError := rtepub.GetExecutionError(err, rLimitErr)
		return &rtepub.Output{StartedOn: startedOn,
			Status:     execError.Status,
			ErrorDescr: execError.Description}
	}
	finishedOn := time.Now()
	elapsedTime := finishedOn.Sub(startedOn)

	err = state.LuaState.CallByParam(lua.P{
		Fn:      state.LuaState.GetGlobal(input.MainFn),
		NRet:    1,
		Protect: true,
	})

	stdout := strings.Join(state.Output.Stdout, "")
	if err != nil {
		mlog.Error("CallByParam error: %v", err)
		i.rLimit.StopMonitoring()
		timer.Stop()
		ErrDoREPL.Incr()
		execError := rtepub.GetExecutionError(err, rLimitErr)
		return &rtepub.Output{StartedOn: startedOn,
			FinishedOn: finishedOn,
			Stdout:     stdout,
			Status:     execError.Status,
			ErrorDescr: execError.Description}
	}

	result := lua.LVAsString(state.LuaState.Get(-1))
	state.LuaState.Pop(1)
	if state.Output.Result != "" {
		result = state.Output.Result
	}

	i.rLimit.StopMonitoring()
	timer.Stop()
	DoREPL.Incr()

	output := &rtepub.Output{
		StartedOn:   startedOn,
		FinishedOn:  finishedOn,
		ElapsedTime: elapsedTime,
		Stdout:      stdout,
		Result:      result,
		Status:      rtepub.SNIPPET_RUN_FINISHED,
		ErrorDescr:  ""}
	mlog.Debug("REPL output: %v", output)
	return output
}

func (i *LuaInterpreter) Terminate() {
	mlog.Debug("Terminating")

	if i.State != nil {
		i.State.Close()
	}
}
