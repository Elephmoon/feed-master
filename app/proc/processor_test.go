package proc

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestProcessor_setDefaults(t *testing.T) {
	type fields struct {
		Conf *Conf
	}
	tests := []struct {
		name   string
		fields fields
		want   *Processor
	}{
		{
			name: "Conf.System.Concurrent == 0",
			fields: fields{
				Conf: &Conf{
					System: struct {
						UpdateInterval time.Duration `yaml:"update"`
						MaxItems       int           `yaml:"max_per_feed"`
						MaxTotal       int           `yaml:"max_total"`
						MaxKeepInDB    int           `yaml:"max_keep"`
						Concurrent     int           `yaml:"concurrent"`
						BaseURL        string        `yaml:"base_url"`
					}{
						UpdateInterval: 1,
						MaxItems:       1,
						MaxTotal:       1,
						MaxKeepInDB:    1,
						Concurrent:     0,
						BaseURL:        "",
					},
				},
			},
			want: &Processor{
				Conf: &Conf{
					System: struct {
						UpdateInterval time.Duration `yaml:"update"`
						MaxItems       int           `yaml:"max_per_feed"`
						MaxTotal       int           `yaml:"max_total"`
						MaxKeepInDB    int           `yaml:"max_keep"`
						Concurrent     int           `yaml:"concurrent"`
						BaseURL        string        `yaml:"base_url"`
					}{
						UpdateInterval: 1,
						MaxItems:       1,
						MaxTotal:       1,
						MaxKeepInDB:    1,
						Concurrent:     8,
						BaseURL:        "",
					},
				},
			},
		},
		{
			name: "Conf.System.MaxItems == 0",
			fields: fields{
				Conf: &Conf{
					System: struct {
						UpdateInterval time.Duration `yaml:"update"`
						MaxItems       int           `yaml:"max_per_feed"`
						MaxTotal       int           `yaml:"max_total"`
						MaxKeepInDB    int           `yaml:"max_keep"`
						Concurrent     int           `yaml:"concurrent"`
						BaseURL        string        `yaml:"base_url"`
					}{
						UpdateInterval: 1,
						MaxItems:       0,
						MaxTotal:       1,
						MaxKeepInDB:    1,
						Concurrent:     1,
						BaseURL:        "",
					},
				},
			},
			want: &Processor{
				Conf: &Conf{
					System: struct {
						UpdateInterval time.Duration `yaml:"update"`
						MaxItems       int           `yaml:"max_per_feed"`
						MaxTotal       int           `yaml:"max_total"`
						MaxKeepInDB    int           `yaml:"max_keep"`
						Concurrent     int           `yaml:"concurrent"`
						BaseURL        string        `yaml:"base_url"`
					}{
						UpdateInterval: 1,
						MaxItems:       5,
						MaxTotal:       1,
						MaxKeepInDB:    1,
						Concurrent:     1,
						BaseURL:        "",
					},
				},
			},
		},
		{
			name: "Conf.System.MaxTotal == 0",
			fields: fields{
				Conf: &Conf{
					System: struct {
						UpdateInterval time.Duration `yaml:"update"`
						MaxItems       int           `yaml:"max_per_feed"`
						MaxTotal       int           `yaml:"max_total"`
						MaxKeepInDB    int           `yaml:"max_keep"`
						Concurrent     int           `yaml:"concurrent"`
						BaseURL        string        `yaml:"base_url"`
					}{
						UpdateInterval: 1,
						MaxItems:       1,
						MaxTotal:       0,
						MaxKeepInDB:    1,
						Concurrent:     1,
						BaseURL:        "",
					},
				},
			},
			want: &Processor{
				Conf: &Conf{
					System: struct {
						UpdateInterval time.Duration `yaml:"update"`
						MaxItems       int           `yaml:"max_per_feed"`
						MaxTotal       int           `yaml:"max_total"`
						MaxKeepInDB    int           `yaml:"max_keep"`
						Concurrent     int           `yaml:"concurrent"`
						BaseURL        string        `yaml:"base_url"`
					}{
						UpdateInterval: 1,
						MaxItems:       1,
						MaxTotal:       100,
						MaxKeepInDB:    1,
						Concurrent:     1,
						BaseURL:        "",
					},
				},
			},
		},
		{
			name: "Conf.System.MaxKeepInDB == 0",
			fields: fields{
				Conf: &Conf{
					System: struct {
						UpdateInterval time.Duration `yaml:"update"`
						MaxItems       int           `yaml:"max_per_feed"`
						MaxTotal       int           `yaml:"max_total"`
						MaxKeepInDB    int           `yaml:"max_keep"`
						Concurrent     int           `yaml:"concurrent"`
						BaseURL        string        `yaml:"base_url"`
					}{
						UpdateInterval: 1,
						MaxItems:       1,
						MaxTotal:       1,
						MaxKeepInDB:    0,
						Concurrent:     1,
						BaseURL:        "",
					},
				},
			},
			want: &Processor{
				Conf: &Conf{
					System: struct {
						UpdateInterval time.Duration `yaml:"update"`
						MaxItems       int           `yaml:"max_per_feed"`
						MaxTotal       int           `yaml:"max_total"`
						MaxKeepInDB    int           `yaml:"max_keep"`
						Concurrent     int           `yaml:"concurrent"`
						BaseURL        string        `yaml:"base_url"`
					}{
						UpdateInterval: 1,
						MaxItems:       1,
						MaxTotal:       1,
						MaxKeepInDB:    5000,
						Concurrent:     1,
						BaseURL:        "",
					},
				},
			},
		},
		{
			name: "Conf.System.UpdateInterval == 0",
			fields: fields{&Conf{
				System: struct {
					UpdateInterval time.Duration `yaml:"update"`
					MaxItems       int           `yaml:"max_per_feed"`
					MaxTotal       int           `yaml:"max_total"`
					MaxKeepInDB    int           `yaml:"max_keep"`
					Concurrent     int           `yaml:"concurrent"`
					BaseURL        string        `yaml:"base_url"`
				}{
					UpdateInterval: 0,
					MaxItems:       1,
					MaxTotal:       1,
					MaxKeepInDB:    1,
					Concurrent:     1,
					BaseURL:        "",
				},
			}},
			want: &Processor{
				Conf: &Conf{
					Feeds: nil,
					System: struct {
						UpdateInterval time.Duration `yaml:"update"`
						MaxItems       int           `yaml:"max_per_feed"`
						MaxTotal       int           `yaml:"max_total"`
						MaxKeepInDB    int           `yaml:"max_keep"`
						Concurrent     int           `yaml:"concurrent"`
						BaseURL        string        `yaml:"base_url"`
					}{
						UpdateInterval: time.Minute * 5,
						MaxItems:       1,
						MaxTotal:       1,
						MaxKeepInDB:    1,
						Concurrent:     1,
						BaseURL:        "",
					},
				},
			},
		},
		{
			name: "all Conf fields not equal 0",
			fields: fields{&Conf{
				Feeds: nil,
				System: struct {
					UpdateInterval time.Duration `yaml:"update"`
					MaxItems       int           `yaml:"max_per_feed"`
					MaxTotal       int           `yaml:"max_total"`
					MaxKeepInDB    int           `yaml:"max_keep"`
					Concurrent     int           `yaml:"concurrent"`
					BaseURL        string        `yaml:"base_url"`
				}{
					UpdateInterval: 1,
					MaxItems:       1,
					MaxTotal:       1,
					MaxKeepInDB:    1,
					Concurrent:     1,
					BaseURL:        "http://example.ex",
				},
			}},
			want: &Processor{
				Conf: &Conf{
					Feeds: nil,
					System: struct {
						UpdateInterval time.Duration `yaml:"update"`
						MaxItems       int           `yaml:"max_per_feed"`
						MaxTotal       int           `yaml:"max_total"`
						MaxKeepInDB    int           `yaml:"max_keep"`
						Concurrent     int           `yaml:"concurrent"`
						BaseURL        string        `yaml:"base_url"`
					}{
						UpdateInterval: 1,
						MaxItems:       1,
						MaxTotal:       1,
						MaxKeepInDB:    1,
						Concurrent:     1,
						BaseURL:        "http://example.ex",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{
				Conf: tt.fields.Conf,
			}
			p.setDefaults()
			assert.Equal(t, tt.want, p)
		})
	}
}
