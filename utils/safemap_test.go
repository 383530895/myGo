// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package utils

import "testing"

var safeMap *SafeMap

func TestNewBeeMap(t *testing.T) {
	safeMap = NewSafeMap()
	if safeMap == nil {
		t.Fatal("expected to return non-nil BeeMap", "got", safeMap)
	}
}

func TestSet(t *testing.T) {
	safeMap = NewSafeMap()
	if ok := safeMap.Set("hehf", 1); !ok {
		t.Error("expected", true, "got", false)
	}
}

func TestCheck(t *testing.T) {
	safeMap = NewSafeMap()
	safeMap.Set("hehf", 2)
	if exists := safeMap.Check("hehf"); !exists {
		t.Error("expected", true, "got", false)
	}
}

func TestGet(t *testing.T) {
	safeMap = NewSafeMap()
	safeMap.Set("hehf", 1)
	if val := safeMap.Get("hehf"); val.(int) != 1 {
		t.Error("expected value", 1, "got", val)
	}
}

func TestDelete(t *testing.T) {
	safeMap = NewSafeMap()
	safeMap.Set("hehf", 2)
	safeMap.Delete("hehf")
	if exists := safeMap.Check("hehf"); exists {
		t.Error("expected element to be deleted")
	}
}

func TestCount(t *testing.T) {
	safeMap = NewSafeMap()
	//safeMap.Set("hehf", 2)
	if count := safeMap.Count(); count != 0 {
		t.Error("expected count to be", 0, "got", count)
	}
}
