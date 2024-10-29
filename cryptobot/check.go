package cryptobot

import (
	"fmt"
	"net/url"
	"strconv"
)

// CreateCheckRequest - params for `CreateCheck`
type CreateCheckRequest struct {
	// Cryptocurrency alphabetic code. Supported assets: “USDT”, “TON”, “BTC”, “ETH”, “LTC”, “BNB”, “TRX” and “USDC” (and “JET” for testnet).
	Asset string `json:"asset"`

	// Amount of the check in float. For example: 125.50
	Amount float64 `json:"amount"`

	// Optional. ID of the user who will be able to activate the check.
	PinToUserID int64 `json:"pin_to_user_id"`

	// Optional. A user with the specified username will be able to activate the check.
	PinToUsername string `json:"pin_to_username"`
}

type createCheckResponse struct {
	response
	Result Check `json:"result"`
}

// CreateCheck - Use this method to create a new check. On success, returns an object of the created Check
func (c *Client) CreateCheck(req CreateCheckRequest) (*Check, error) {
	responseBodyReader, err := c.request("createCheck", func(q url.Values) url.Values {
		q.Add("asset", req.Asset)
		q.Add("amount", strconv.FormatFloat(req.Amount, 'f', -1, 64))

		if req.PinToUserID != 0 {
			q.Add("pin_to_user_id", strconv.Itoa(int(req.PinToUserID)))
		}

		if req.PinToUsername != "" {
			q.Add("pin_to_username", req.PinToUsername)
		}

		return q
	})

	if err != nil {
		return nil, err
	}
	defer responseBodyReader.Close()

	var response createCheckResponse
	if err := c.decodeResponse(responseBodyReader, &response); err != nil {
		return nil, err
	}

	if response.Ok {
		return &response.Result, nil
	}

	return nil, fmt.Errorf("createCheck request error: code - %v, name - %s", response.Error.Code, response.Error.Name)
}

// DeleteCheckRequest - params for `DeleteCheck`
type DeleteCheckRequest struct {
	// Check ID to be deleted.
	CheckID int64 `json:"check_id"`
}

type deleteCheckResponse struct {
	response
	Result bool `json:"result"`
}

// DeleteCheck - Use this method to delete checks created by your app. Returns `true` on success.
func (c *Client) DeleteCheck(req DeleteCheckRequest) (bool, error) {
	responseBodyReader, err := c.request("deleteCheck", func(q url.Values) url.Values {
		q.Add("check_id", strconv.Itoa(int(req.CheckID)))

		return q
	})

	if err != nil {
		return false, err
	}
	defer responseBodyReader.Close()

	var response deleteCheckResponse
	if err := c.decodeResponse(responseBodyReader, &response); err != nil {
		return false, err
	}

	if response.Ok {
		return response.Result, nil
	}

	return false, fmt.Errorf("deleteCheck request error: code - %v, name - %s", response.Error.Code, response.Error.Name)
}

// GetChecksRequest - params for `GetChecks`
type GetChecksRequest struct {
	// Optional. Cryptocurrency alphabetic code. Supported assets: “USDT”, “TON”, “BTC”, “ETH”, “LTC”, “BNB”, “TRX” and “USDC” (and “JET” for testnet). Defaults to all currencies.
	Asset string `json:"asset"`

	// Optional. List of check IDs separated by comma.
	CheckIDs []int64 `json:"check_ids"`

	// Optional. Status of check to be returned. Available statuses: “active” and “activated”. Defaults to all statuses.
	Status string `json:"status"`

	// Optional. Offset needed to return a specific subset of check. Defaults to 0.
	Offset int64 `json:"offset"`

	// Optional. Number of check to be returned. Values between 1-1000 are accepted. Defaults to 100.
	Count int64 `json:"count"`
}

type getChecksResponse struct {
	response
	Result []Check `json:"result"`
}

// GetChecks - Use this method to get checks created by your app. On success, returns array of `Check“.
func (c *Client) GetChecks(req GetChecksRequest) ([]Check, error) {
	responseBodyReader, err := c.request("getChecks", func(q url.Values) url.Values {
		if req.Asset != "" {
			q.Add("asset", req.Asset)
		}

		if len(req.CheckIDs) != 0 {
			q.Add("check_ids", idsToString(req.CheckIDs))
		}

		if req.Status != "" {
			q.Add("status", req.Status)
		}

		if req.Offset != 0 {
			q.Add("offset", strconv.Itoa(int(req.Offset)))
		}

		if req.Count != 0 {
			q.Add("count", strconv.Itoa(int(req.Count)))
		}

		return q
	})

	if err != nil {
		return nil, err
	}
	defer responseBodyReader.Close()

	var response getChecksResponse
	if err := c.decodeResponse(responseBodyReader, &response); err != nil {
		return nil, err
	}

	if response.Ok {
		return response.Result, nil
	}

	return nil, fmt.Errorf("getChecks request error: code - %v, name - %s", response.Error.Code, response.Error.Name)
}

func idsToString(IDs []int64) string {
	if len(IDs) == 0 {
		return ""
	}

	out := strconv.Itoa(int(IDs[0]))

	for _, ID := range IDs {
		out += "," + strconv.Itoa(int(ID))
	}

	return out
}
