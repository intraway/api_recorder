package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func validate_aaa(data []*EnhancedRequest, n int, t *testing.T) {
	if n <= 0 || n >= 3 {
		if len(data) != 3 {
			t.Errorf("Len '/aaa' entry != 3. Got: %v", len(data))
		}
	} else {
		if len(data) != n {
			t.Errorf("Len '/aaa' entry != %v. Got: %v", n, len(data))
		}
	}
}

func validate_bbbccc(data []*EnhancedRequest, t *testing.T) {
	if len(data) != 2 {
		t.Errorf("Len '/bbb/ccc' entry != 2. Got: %v", len(data))
	}
}

func validate_ddd(data []*EnhancedRequest, t *testing.T) {
	if len(data) != 1 {
		t.Errorf("Len '/ddd' entry != 1. Got: %v", len(data))
	}
}

func TestShowAndReset(t *testing.T) {
	config := DefaultConfig()
	base_address := fmt.Sprintf("http://%v:%v", config.Host, config.Port)

	rm := NewRequestsManager(DefaultConfig())
	go rm.Run()
	// Give some time to start
	time.Sleep(time.Millisecond * 10)

	client := http.Client{}
	endpoint1 := base_address + "/aaa"
	endpoint2 := base_address + "/bbb/ccc"
	endpoint3 := base_address + "/ddd"

	client.Get(endpoint1)
	client.Get(endpoint1 + "?q1=10&q2=20")
	client.Get(endpoint1 + "?q1=30&q2=40")
	client.Get(endpoint2)
	client.Get(endpoint2)
	e3body := []byte(`{"abc":123}`)
	client.Post(endpoint3, "application/json", bytes.NewBuffer(e3body))

	// GET ALL
	resp, err := client.Get(base_address + "/" + config.ShowURL)
	if err != nil {
		t.Errorf("GET error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Server responded with status '%v'", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Errorf("Body read error: %v", err)
	}
	r := make(map[string][]*EnhancedRequest)

	err = json.Unmarshal(body, &r)
	if err != nil {
		t.Errorf("Error unmarshaling: %v", err)
	} else {
		//t.Errorf("%+v", string(body))
		if aaa, ok := r["/aaa"]; !ok {
			t.Errorf("No '/aaa' entry")
		} else {
			validate_aaa(aaa, 0, t)
		}

		if bbbccc, ok := r["/bbb/ccc"]; !ok {
			t.Errorf("No '/bbb/ccc' entry")
		} else {
			validate_bbbccc(bbbccc, t)
		}

		if ddd, ok := r["/ddd"]; !ok {
			t.Errorf("No '/ddd' entry")
		} else {
			validate_ddd(ddd, t)
		}
	}

	// Individual GET /aaa
	resp, err = client.Get(base_address + "/" + config.ShowURL + "?url=/aaa")
	if err != nil {
		t.Errorf("GET error: %v", err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Errorf("Body read error: %v", err)
	}
	r = make(map[string][]*EnhancedRequest)

	err = json.Unmarshal(body, &r)
	if err != nil {
		t.Errorf("Error unmarshaling: %v", err)
	} else {
		if len(r) != 1 {
			t.Errorf("Only one result was expected. Got %v", len(r))
		} else {
			if aaa, ok := r["/aaa"]; !ok {
				t.Errorf("No '/aaa' entry")
			} else {
				validate_aaa(aaa, 0, t)
			}
		}
	}

	// Individual GET /aaa with big n
	resp, err = client.Get(base_address + "/" + config.ShowURL + "?url=/aaa&n=20")
	if err != nil {
		t.Errorf("GET error: %v", err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Errorf("Body read error: %v", err)
	}
	r = make(map[string][]*EnhancedRequest)

	err = json.Unmarshal(body, &r)
	if err != nil {
		t.Errorf("Error unmarshaling: %v", err)
	} else {
		if len(r) != 1 {
			t.Errorf("Only one result was expected. Got %v", len(r))
		} else {
			if aaa, ok := r["/aaa"]; !ok {
				t.Errorf("No '/aaa' entry")
			} else {
				validate_aaa(aaa, 20, t)
			}
		}
	}

	// Individual GET /aaa with n = 1
	resp, err = client.Get(base_address + "/" + config.ShowURL + "?url=/aaa&n=1")
	if err != nil {
		t.Errorf("GET error: %v", err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Errorf("Body read error: %v", err)
	}
	r = make(map[string][]*EnhancedRequest)

	err = json.Unmarshal(body, &r)
	if err != nil {
		t.Errorf("Error unmarshaling: %v", err)
	} else {
		if len(r) != 1 {
			t.Errorf("Only one result was expected. Got %v", len(r))
		} else {
			if aaa, ok := r["/aaa"]; !ok {
				t.Errorf("No '/aaa' entry")
			} else {
				validate_aaa(aaa, 1, t)
			}
		}
	}

	// Individual GET /bbb/ccc
	resp, err = client.Get(base_address + "/" + config.ShowURL + "?url=/bbb/ccc")
	if err != nil {
		t.Errorf("GET error: %v", err)
	}

	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Errorf("Body read error: %v", err)
	}
	r = make(map[string][]*EnhancedRequest)

	err = json.Unmarshal(body, &r)
	if err != nil {
		t.Errorf("Error unmarshaling: %v", err)
	} else {
		if len(r) != 1 {
			t.Errorf("Only one result was expected. Got %v", len(r))
		} else {
			if bbbccc, ok := r["/bbb/ccc"]; !ok {
				t.Errorf("No '/bbb/ccc' entry")
			} else {
				validate_bbbccc(bbbccc, t)
			}
		}
	}

	// Reset /aaa
	client.Get(base_address + "/" + config.ResetURL + "?url=/aaa")

	// GET all
	resp, err = client.Get(base_address + "/" + config.ShowURL)
	if err != nil {
		t.Errorf("GET error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Server responded with status '%v'", resp.Status)
	}

	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Errorf("Body read error: %v", err)
	}
	r = make(map[string][]*EnhancedRequest)

	err = json.Unmarshal(body, &r)
	if err != nil {
		t.Errorf("Error unmarshaling: %v", err)
	} else {
		if _, ok := r["/aaa"]; ok {
			t.Errorf("'/aaa' shouldn't be present")
		}

		if bbbccc, ok := r["/bbb/ccc"]; !ok {
			t.Errorf("No '/bbb/ccc' entry")
		} else {
			validate_bbbccc(bbbccc, t)
		}

		if ddd, ok := r["/ddd"]; !ok {
			t.Errorf("No '/ddd' entry")
		} else {
			validate_ddd(ddd, t)
		}
	}

	// Reset all
	client.Get(base_address + "/" + config.ResetURL)

	// GET all
	resp, err = client.Get(base_address + "/" + config.ShowURL)
	if err != nil {
		t.Errorf("GET error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Server responded with status '%v'", resp.Status)
	}

	body, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Errorf("Body read error: %v", err)
	}
	r = make(map[string][]*EnhancedRequest)

	err = json.Unmarshal(body, &r)
	if err != nil {
		t.Errorf("Error unmarshaling: %v", err)
	} else {
		if len(r) != 0 {
			t.Errorf("No elements expected. Got %v", len(r))
		}
	}
}
