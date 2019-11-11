// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: assetupdatefeedproduceroperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/blocktree/whitecoin-adapter/libs/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *AssetUpdateFeedProducersOperation) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *AssetUpdateFeedProducersOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "asset_to_update":`)

	{

		obj, err = j.AssetToUpdate.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"extensions":`)

	{

		obj, err = j.Extensions.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"issuer":`)

	{

		obj, err = j.Issuer.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"new_feed_producers":`)
	if j.NewFeedProducers != nil {
		buf.WriteString(`[`)
		for i, v := range j.NewFeedProducers {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				obj, err = v.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte(',')
	if j.Fee != nil {
		if true {
			/* Struct fall back. type=types.AssetAmount kind=struct */
			buf.WriteString(`"fee":`)
			err = buf.Encode(j.Fee)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAssetUpdateFeedProducersOperationbase = iota
	ffjtAssetUpdateFeedProducersOperationnosuchkey

	ffjtAssetUpdateFeedProducersOperationAssetToUpdate

	ffjtAssetUpdateFeedProducersOperationExtensions

	ffjtAssetUpdateFeedProducersOperationIssuer

	ffjtAssetUpdateFeedProducersOperationNewFeedProducers

	ffjtAssetUpdateFeedProducersOperationFee
)

var ffjKeyAssetUpdateFeedProducersOperationAssetToUpdate = []byte("asset_to_update")

var ffjKeyAssetUpdateFeedProducersOperationExtensions = []byte("extensions")

var ffjKeyAssetUpdateFeedProducersOperationIssuer = []byte("issuer")

var ffjKeyAssetUpdateFeedProducersOperationNewFeedProducers = []byte("new_feed_producers")

var ffjKeyAssetUpdateFeedProducersOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AssetUpdateFeedProducersOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AssetUpdateFeedProducersOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAssetUpdateFeedProducersOperationbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtAssetUpdateFeedProducersOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyAssetUpdateFeedProducersOperationAssetToUpdate, kn) {
						currentKey = ffjtAssetUpdateFeedProducersOperationAssetToUpdate
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyAssetUpdateFeedProducersOperationExtensions, kn) {
						currentKey = ffjtAssetUpdateFeedProducersOperationExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyAssetUpdateFeedProducersOperationFee, kn) {
						currentKey = ffjtAssetUpdateFeedProducersOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyAssetUpdateFeedProducersOperationIssuer, kn) {
						currentKey = ffjtAssetUpdateFeedProducersOperationIssuer
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyAssetUpdateFeedProducersOperationNewFeedProducers, kn) {
						currentKey = ffjtAssetUpdateFeedProducersOperationNewFeedProducers
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyAssetUpdateFeedProducersOperationFee, kn) {
					currentKey = ffjtAssetUpdateFeedProducersOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateFeedProducersOperationNewFeedProducers, kn) {
					currentKey = ffjtAssetUpdateFeedProducersOperationNewFeedProducers
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateFeedProducersOperationIssuer, kn) {
					currentKey = ffjtAssetUpdateFeedProducersOperationIssuer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateFeedProducersOperationExtensions, kn) {
					currentKey = ffjtAssetUpdateFeedProducersOperationExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateFeedProducersOperationAssetToUpdate, kn) {
					currentKey = ffjtAssetUpdateFeedProducersOperationAssetToUpdate
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAssetUpdateFeedProducersOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtAssetUpdateFeedProducersOperationAssetToUpdate:
					goto handle_AssetToUpdate

				case ffjtAssetUpdateFeedProducersOperationExtensions:
					goto handle_Extensions

				case ffjtAssetUpdateFeedProducersOperationIssuer:
					goto handle_Issuer

				case ffjtAssetUpdateFeedProducersOperationNewFeedProducers:
					goto handle_NewFeedProducers

				case ffjtAssetUpdateFeedProducersOperationFee:
					goto handle_Fee

				case ffjtAssetUpdateFeedProducersOperationnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_AssetToUpdate:

	/* handler: j.AssetToUpdate type=types.AssetID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.AssetToUpdate.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Extensions:

	/* handler: j.Extensions type=types.Extensions kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Extensions.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Issuer:

	/* handler: j.Issuer type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Issuer.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NewFeedProducers:

	/* handler: j.NewFeedProducers type=types.AccountIDs kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for AccountIDs", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.NewFeedProducers = nil
		} else {

			j.NewFeedProducers = []types.AccountID{}

			wantVal := true

			for {

				var tmpJNewFeedProducers types.AccountID

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJNewFeedProducers type=types.AccountID kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJNewFeedProducers.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.NewFeedProducers = append(j.NewFeedProducers, tmpJNewFeedProducers)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Fee)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
