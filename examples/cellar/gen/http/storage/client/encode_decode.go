// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// storage HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/goa/examples/cellar/design -o
// $(GOPATH)/src/goa.design/goa/examples/cellar

package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"

	storage "goa.design/goa/examples/cellar/gen/storage"
	goahttp "goa.design/goa/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "storage" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListStoragePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("storage", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListResponse returns a decoder for responses returned by the storage
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("storage", "list", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, fmt.Errorf("invalid response: %s", err)
			}

			return NewListStoredBottleCollectionOK(body), nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("account", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildShowRequest instantiates a HTTP request object with method and path set
// to call the "storage" service "show" endpoint
func (c *Client) BuildShowRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*storage.ShowPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("storage", "show", "*storage.ShowPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowStoragePath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("storage", "show", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeShowRequest returns an encoder for requests sent to the storage show
// server.
func EncodeShowRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*storage.ShowPayload)
		if !ok {
			return goahttp.ErrInvalidType("storage", "show", "*storage.ShowPayload", v)
		}
		values := req.URL.Query()
		if p.View != nil {
			values.Add("view", *p.View)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeShowResponse returns a decoder for responses returned by the storage
// show endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeShowResponse may return the following errors:
//	- "not_found" (type *storage.NotFound): http.StatusNotFound
//	- error: internal error
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ShowResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("storage", "show", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, fmt.Errorf("invalid response: %s", err)
			}

			return NewShowStoredBottleOK(&body), nil
		case http.StatusNotFound:
			var (
				body ShowNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("storage", "show", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, fmt.Errorf("invalid response: %s", err)
			}

			return nil, NewShowNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("account", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildAddRequest instantiates a HTTP request object with method and path set
// to call the "storage" service "add" endpoint
func (c *Client) BuildAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddStoragePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("storage", "add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddRequest returns an encoder for requests sent to the storage add
// server.
func EncodeAddRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*storage.Bottle)
		if !ok {
			return goahttp.ErrInvalidType("storage", "add", "*storage.Bottle", v)
		}
		body := NewAddRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("storage", "add", err)
		}
		return nil
	}
}

// DecodeAddResponse returns a decoder for responses returned by the storage
// add endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("storage", "add", err)
			}

			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("account", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildRemoveRequest instantiates a HTTP request object with method and path
// set to call the "storage" service "remove" endpoint
func (c *Client) BuildRemoveRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*storage.RemovePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("storage", "remove", "*storage.RemovePayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RemoveStoragePath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("storage", "remove", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeRemoveResponse returns a decoder for responses returned by the storage
// remove endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeRemoveResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("account", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildRateRequest instantiates a HTTP request object with method and path set
// to call the "storage" service "rate" endpoint
func (c *Client) BuildRateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RateStoragePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("storage", "rate", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRateRequest returns an encoder for requests sent to the storage rate
// server.
func EncodeRateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(map[uint32][]string)
		if !ok {
			return goahttp.ErrInvalidType("storage", "rate", "map[uint32][]string", v)
		}
		values := req.URL.Query()
		for key, value := range p {
			keyStr := strconv.FormatUint(uint64(key), 10)
			for _, val := range value {
				valStr := val
				values.Add(keyStr, valStr)
			}
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeRateResponse returns a decoder for responses returned by the storage
// rate endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeRateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("account", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildMultiAddRequest instantiates a HTTP request object with method and path
// set to call the "storage" service "multi_add" endpoint
func (c *Client) BuildMultiAddRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MultiAddStoragePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("storage", "multi_add", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeMultiAddRequest returns an encoder for requests sent to the storage
// multi_add server.
func EncodeMultiAddRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.([]*storage.Bottle)
		if !ok {
			return goahttp.ErrInvalidType("storage", "multi_add", "[]*storage.Bottle", v)
		}
		if err := encoder(req).Encode(p); err != nil {
			return goahttp.ErrEncodingError("storage", "multi_add", err)
		}
		return nil
	}
}

// NewStorageMultiAddEncoder returns an encoder to encode the multipart request
// for the "storage" service "multi_add" endpoint.
func NewStorageMultiAddEncoder(encoderFn StorageMultiAddEncoderFunc) func(r *http.Request) goahttp.Encoder {
	return func(r *http.Request) goahttp.Encoder {
		body := &bytes.Buffer{}
		mw := multipart.NewWriter(body)
		return goahttp.EncodingFunc(func(v interface{}) error {
			p := v.([]*storage.Bottle)
			if err := encoderFn(mw, p); err != nil {
				return err
			}
			r.Body = ioutil.NopCloser(body)
			r.Header.Set("Content-Type", mw.FormDataContentType())
			return mw.Close()
		})
	}
}

// DecodeMultiAddResponse returns a decoder for responses returned by the
// storage multi_add endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeMultiAddResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body []string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("storage", "multi_add", err)
			}

			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("account", "create", resp.StatusCode, string(body))
		}
	}
}

// BuildMultiUpdateRequest instantiates a HTTP request object with method and
// path set to call the "storage" service "multi_update" endpoint
func (c *Client) BuildMultiUpdateRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MultiUpdateStoragePath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("storage", "multi_update", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeMultiUpdateRequest returns an encoder for requests sent to the storage
// multi_update server.
func EncodeMultiUpdateRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*storage.MultiUpdatePayload)
		if !ok {
			return goahttp.ErrInvalidType("storage", "multi_update", "*storage.MultiUpdatePayload", v)
		}
		values := req.URL.Query()
		for _, value := range p.Ids {
			values.Add("ids", value)
		}
		req.URL.RawQuery = values.Encode()
		if err := encoder(req).Encode(p); err != nil {
			return goahttp.ErrEncodingError("storage", "multi_update", err)
		}
		return nil
	}
}

// NewStorageMultiUpdateEncoder returns an encoder to encode the multipart
// request for the "storage" service "multi_update" endpoint.
func NewStorageMultiUpdateEncoder(encoderFn StorageMultiUpdateEncoderFunc) func(r *http.Request) goahttp.Encoder {
	return func(r *http.Request) goahttp.Encoder {
		body := &bytes.Buffer{}
		mw := multipart.NewWriter(body)
		return goahttp.EncodingFunc(func(v interface{}) error {
			p := v.(*storage.MultiUpdatePayload)
			if err := encoderFn(mw, p); err != nil {
				return err
			}
			r.Body = ioutil.NopCloser(body)
			r.Header.Set("Content-Type", mw.FormDataContentType())
			return mw.Close()
		})
	}
}

// DecodeMultiUpdateResponse returns a decoder for responses returned by the
// storage multi_update endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeMultiUpdateResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("account", "create", resp.StatusCode, string(body))
		}
	}
}

// unmarshalWineryResponseBodyToWinery builds a value of type *storage.Winery
// from a value of type *WineryResponseBody.
func unmarshalWineryResponseBodyToWinery(v *WineryResponseBody) *storage.Winery {
	res := &storage.Winery{
		Name:    *v.Name,
		Region:  *v.Region,
		Country: *v.Country,
		URL:     v.URL,
	}

	return res
}

// unmarshalWineryToWinery builds a value of type *storage.Winery from a value
// of type *Winery.
func unmarshalWineryToWinery(v *Winery) *storage.Winery {
	res := &storage.Winery{
		Name:    *v.Name,
		Region:  *v.Region,
		Country: *v.Country,
		URL:     v.URL,
	}

	return res
}

// marshalWineryToWineryRequestBody builds a value of type *WineryRequestBody
// from a value of type *storage.Winery.
func marshalWineryToWineryRequestBody(v *storage.Winery) *WineryRequestBody {
	res := &WineryRequestBody{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
		URL:     v.URL,
	}

	return res
}

// marshalWineryRequestBodyToWinery builds a value of type *storage.Winery from
// a value of type *WineryRequestBody.
func marshalWineryRequestBodyToWinery(v *WineryRequestBody) *storage.Winery {
	res := &storage.Winery{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
		URL:     v.URL,
	}

	return res
}
