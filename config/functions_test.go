package config

import(
    "testing"
)

func TestConfig(t *testing.T) {
    ReadFromFile("../config.json")
    if Database != "sttest.sqlt" {
        t.Errorf("Incorrect database name")
    }
}