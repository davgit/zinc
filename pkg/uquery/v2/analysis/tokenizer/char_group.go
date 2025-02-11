package tokenizer

import (
	"github.com/blugelabs/bluge/analysis"

	zinctokenizer "github.com/zinclabs/zinc/pkg/bluge/analysis/tokenizer"
	"github.com/zinclabs/zinc/pkg/errors"
	"github.com/zinclabs/zinc/pkg/zutils"
)

func NewCharGroupTokenizer(options interface{}) (analysis.Tokenizer, error) {
	chars, err := zutils.GetStringSliceFromMap(options, "tokenize_on_chars")
	if err != nil {
		return nil, errors.New(errors.ErrorTypeParsingException, "[tokenizer] char_group option [tokenize_on_chars] should be an array of string")
	}

	return zinctokenizer.NewCharGroupTokenizer(chars), nil
}
