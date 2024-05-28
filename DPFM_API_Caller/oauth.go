package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-instagram-auth-key-requests-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-instagram-auth-key-requests-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-instagram-auth-key-requests-rmq-kube/config"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
	"net/url"
)

func (c *DPFMAPICaller) InstagramAuthKey(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	log *logger.Logger,
	conf *config.Conf,
) *[]dpfm_api_output_formatter.InstagramAuthKey {
	var instagramAuthKey []dpfm_api_output_formatter.InstagramAuthKey

	inputURL := input.InstagramAuthKey.URL

	parsedURL, err := url.Parse(inputURL)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	code := parsedURL.Query().Get("code")

	if code == "" {
		*errs = append(*errs, xerrors.Errorf("URL does not contain Code"))
		return nil
	}

	generatedInstagramAuthKey := conf.OAuth.GenerateOAuthTokenRequestURL()

	instagramAuthKey = append(instagramAuthKey, dpfm_api_output_formatter.InstagramAuthKey{
		URL:  generatedInstagramAuthKey,
		Code: code,
	})

	return &instagramAuthKey
}
