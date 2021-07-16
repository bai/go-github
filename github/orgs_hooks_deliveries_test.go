// Copyright 2021 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOrganizationsService_ListHookDeliveries(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/hooks/1/deliveries", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{"cursor": "v1_12077215967"})
		fmt.Fprint(w, `[{"id":1}, {"id":2}]`)
	})

	opt := &ListCursorOptions{Cursor: "v1_12077215967"}

	ctx := context.Background()
	hooks, _, err := client.Organizations.ListHookDeliveries(ctx, "o", 1, opt)
	if err != nil {
		t.Errorf("Organizations.ListHookDeliveries returned error: %v", err)
	}

	want := []*HookDelivery{{ID: Int64(1)}, {ID: Int64(2)}}
	if d := cmp.Diff(hooks, want); d != "" {
		t.Errorf("Organizations.ListHooks want (-), got (+):\n%s", d)
	}

	const methodName = "ListHookDeliveries"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Organizations.ListHookDeliveries(ctx, "\n", -1, opt)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Organizations.ListHookDeliveries(ctx, "o", 1, opt)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestOrganizationsService_ListHookDeliveries_invalidOwner(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	ctx := context.Background()
	_, _, err := client.Organizations.ListHookDeliveries(ctx, "%", 1, nil)
	testURLParseError(t, err)
}

func TestOrganizationsService_GetHookDelivery(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/orgs/o/hooks/1/deliveries/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1}`)
	})

	ctx := context.Background()
	hook, _, err := client.Organizations.GetHookDelivery(ctx, "o", 1, 1)
	if err != nil {
		t.Errorf("Organizations.GetHookDelivery returned error: %v", err)
	}

	want := &HookDelivery{ID: Int64(1)}
	if !cmp.Equal(hook, want) {
		t.Errorf("Organizations.GetHookDelivery returned %+v, want %+v", hook, want)
	}

	const methodName = "GetHookDelivery"
	testBadOptions(t, methodName, func() (err error) {
		_, _, err = client.Organizations.GetHookDelivery(ctx, "\n", -1, -1)
		return err
	})

	testNewRequestAndDoFailure(t, methodName, client, func() (*Response, error) {
		got, resp, err := client.Organizations.GetHookDelivery(ctx, "o", 1, 1)
		if got != nil {
			t.Errorf("testNewRequestAndDoFailure %v = %#v, want nil", methodName, got)
		}
		return resp, err
	})
}

func TestOrganizationsService_GetHookDelivery_invalidOwner(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	ctx := context.Background()
	_, _, err := client.Organizations.GetHookDelivery(ctx, "%", 1, 1)
	testURLParseError(t, err)
}
