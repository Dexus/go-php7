// Copyright 2016 Alexander Palaistras. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package engine

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestContextNew(t *testing.T) {
	Initialize()
	c := &Context{}
	err := RequestStartup(c)
	if err != nil {
		t.Fatalf("NewContext(): %s", err)
	}

	if c.context == nil || c.ResponseWriter == nil {
		t.Fatalf("NewContext(): Struct fields are `nil` but no error returned")
	}

	RequestShutdown(c)
}

var execTests = []struct {
	name     string
	script   string
	expected string
}{
	{
		"helloworld.php",

		`<?php
		$a = 'Hello';
		$b = 'World';
		echo $a.' '.$b;`,

		"Hello World",
	},
}

func TestContextExec(t *testing.T) {
	Initialize()
	var w bytes.Buffer

	c := &Context{
		Output: &w,
	}
	RequestStartup(c)

	for _, tt := range execTests {
		script, err := NewScript(tt.name, tt.script)
		if err != nil {
			t.Errorf("Could not create temporary file '%s' for testing: %s", tt.name, err)
			continue
		}

		if err := c.Exec(script.Name()); err != nil {
			t.Errorf("Context.Exec('%s'): Execution failed: %s", tt.name, err)
			continue
		}

		actual := w.String()
		w.Reset()

		if actual != tt.expected {
			t.Errorf("Context.Exec('%s'): Expected `%s', actual `%s'", tt.name, tt.expected, actual)
		}

		script.Remove()
	}

	RequestShutdown(c)
}

var evalTests = []struct {
	script string
	output string
	value  interface{}
}{
	{
		"echo 'Hello World';",
		"Hello World",
		nil,
	},
	{
		"$i = 10; $d = 20; return $i + $d;",
		"",
		int64(30),
	},
}

func TestContextEval(t *testing.T) {
	Initialize()
	var w bytes.Buffer

	c := &Context{
		Output: &w,
	}
	RequestStartup(c)

	for _, tt := range evalTests {
		val, err := c.Eval(tt.script)
		if err != nil {
			t.Errorf("Context.Eval('%s'): %s", tt.script, err)
			continue
		}

		output := w.String()
		w.Reset()

		if output != tt.output {
			t.Errorf("Context.Eval('%s'): Expected output '%s', actual '%s'", tt.script, tt.output, output)
		}

		result := ToInterface(val)

		if reflect.DeepEqual(result, tt.value) == false {
			t.Errorf("Context.Eval('%s'): Expected value '%#v', actual '%#v'", tt.script, tt.value, result)
		}

		DestroyValue(val)
	}

	RequestShutdown(c)
}

var logTests = []struct {
	script   string
	expected string
}{
	{
		"$a = 10; $a + $b;",
		"Undefined variable: b in gophp-engine on line 1",
	},
	{
		"strlen();",
		"strlen() expects exactly 1 parameter, 0 given in gophp-engine on line 1",
	},
	{
		"trigger_error('Test Error');",
		"Test Error in gophp-engine on line 1",
	},
}

func TestContextLog(t *testing.T) {
	Initialize()
	var w bytes.Buffer

	c := &Context{
		Log: &w,
	}
	RequestStartup(c)

	for _, tt := range logTests {
		if _, err := c.Eval(tt.script); err != nil {
			t.Errorf("Context.Eval('%s'): %s", tt.script, err)
			continue
		}

		actual := w.String()
		w.Reset()

		if strings.Contains(actual, tt.expected) == false {
			t.Errorf("Context.Eval('%s'): expected '%s', actual '%s'", tt.script, tt.expected, actual)
		}
	}

	RequestShutdown(c)
}

var bindTests = []struct {
	value    interface{}
	expected string
}{
	{
		42,
		"i:42;",
	},
	{
		3.14159,
		"d:3.1415899999999999;",
	},
	{
		true,
		"b:1;",
	},
	{
		"Such bind",
		`s:9:"Such bind";`,
	},
	{
		[]string{"this", "that"},
		`a:2:{i:0;s:4:"this";i:1;s:4:"that";}`,
	},
	{
		[][]int{{1, 2}, {3}},
		`a:2:{i:0;a:2:{i:0;i:1;i:1;i:2;}i:1;a:1:{i:0;i:3;}}`,
	},
	{
		map[string]interface{}{"hello": []string{"hello", "!"}},
		`a:1:{s:5:"hello";a:2:{i:0;s:5:"hello";i:1;s:1:"!";}}`,
	},
	{
		struct {
			I int
			C string
			F struct {
				G bool
			}
			h bool
		}{3, "test", struct {
			G bool
		}{false}, true},
		`O:8:"stdClass":3:{s:1:"I";i:3;s:1:"C";s:4:"test";s:1:"F";O:8:"stdClass":1:{s:1:"G";b:0;}}`,
	},
}

func TestContextBind(t *testing.T) {
	Initialize()
	var w bytes.Buffer

	c := &Context{
		Output: &w,
	}
	RequestStartup(c)

	for i, tt := range bindTests {
		if err := c.Bind(fmt.Sprintf("t%d", i), tt.value); err != nil {
			t.Errorf("Context.Bind('%v'): %s", tt.value, err)
			continue
		}

		if _, err := c.Eval(fmt.Sprintf("echo serialize($t%d);", i)); err != nil {
			t.Errorf("Context.Eval(): %s", err)
			continue
		}

		actual := w.String()
		w.Reset()

		if actual != tt.expected {
			t.Errorf("Context.Bind('%v'): expected '%s', actual '%s'", tt.value, tt.expected, actual)
		}
	}

	RequestShutdown(c)
}

func TestContextDestroy(t *testing.T) {
	Initialize()
	c := &Context{}
	RequestStartup(c)
	RequestShutdown(c)

	if c.context != nil {
		t.Errorf("Context.Destroy(): Did not set internal fields to `nil`")
	}

	// Attempting to destroy a context twice should be a no-op.
	RequestShutdown(c)
}

