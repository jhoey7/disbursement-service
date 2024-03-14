package models

import "net/url"

// Params struct
type Params struct {
	Key   string
	Value string
}

// HTTPParams struct
type HTTPParams struct {
	Params []Params
}

// BuildParam build http param
func (hp *HTTPParams) BuildParam(k string, v string) {
	hp.Params = append(hp.Params, Params{k, v})
}

// URLValues get url values
func (hp HTTPParams) URLValues() url.Values {
	qp := url.Values{}
	for _, p := range hp.Params {
		if p.Key != "" && p.Value != "" {
			qp.Add(p.Key, p.Value)
		}
	}

	return qp
}
