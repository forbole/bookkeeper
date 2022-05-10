package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	cosmostypes "github.com/forbole/bookkeeper/module/cosmos/types"
	"google.golang.org/grpc/metadata"
)

// GetHeightRequestContext adds the height to the context for queries
func GetHeightRequestContext(context context.Context, height int64) context.Context {
	return metadata.AppendToOutgoingContext(
		context,
		grpctypes.GRPCBlockHeightHeader,
		strconv.FormatInt(height, 10),
	)
}

// GetHeightByDate get height for the cloest time stamp within 10 seconds, log2 complexity
// If the height do not exist, will return the Lowest Height
func GetHeightByDate(t time.Time, lcd string) (int, error) {
	query := fmt.Sprintf(`%s/blocks/latest`, lcd)
	fmt.Println(query)
	resp, err := http.Get(query)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("Fail to get tx from rpc:Status code:%d", resp.StatusCode)
	}

	defer resp.Body.Close()

	bz, err := io.ReadAll(resp.Body)

	var blockRes cosmostypes.Block
	err = json.Unmarshal(bz, &blockRes)
	if err != nil {
		return 0, fmt.Errorf("Fail to marshal:%s", err)
	}

	header := blockRes.Block.Header

	latestHeight, err := strconv.Atoi(header.Height)
	if err != nil {
		return 0, err
	}
	left := latestHeight
	newT := header.Time
	right := 0
	middle := (left + right) / 2
	if err != nil {
		return 0, err
	}
	fmt.Println(t)

	for !((t.After(newT) && t.Sub(newT) < (time.Hour*23)) ||
		(t.Before(newT) && newT.Sub(t) < (time.Hour*23))) {
		middle = (left + right) / 2

		if middle == 1 {
			// 1 is the earliest block time
			return 1, nil
		}

		T, err := GetTimeByHeight(middle, lcd)

		if err != nil && strings.Contains(err.Error(), "is not available, lowest height is ") {
			lastIndex := strings.LastIndex(err.Error(), " ")
			lowHeightString := err.Error()[lastIndex+1:]
			lowestHeight, err := strconv.Atoi(lowHeightString)
			if err != nil {
				return 0, err
			}

			// Check if requested t is out of chain scope
			lowestTime, err := GetTimeByHeight(lowestHeight, lcd)
			if err != nil {
				return 0, err
			}
			if t.Before(*lowestTime) {
				//return 0,fmt.Errorf("Request time is out of scope: eariest time: %s, request time: %s",
				//*lowestTime,t)
				return lowestHeight, nil
			}

			right = lowestHeight
			middle = (left + right) / 2
			continue
		}
		if err != nil {
			return 0, err
		}

		newT = *T
		if newT.After(t) {
			left = middle + 1
		} else if newT.Before(t) {
			right = middle - 1
		}
	}
	return middle, nil
}

func GetTimeByHeight(height int, lcd string) (*time.Time, error) {
	query := fmt.Sprintf(`%s/blocks/%d`, lcd, height)
	fmt.Println(query)
	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bz, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		type Err struct {
			E string `json:"error"`
		}

		var e Err
		err = json.Unmarshal(bz, &e)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(e.E)
	}

	bz, err := io.ReadAll(resp.Body)

	var blockRes cosmostypes.Block
	err = json.Unmarshal(bz, &blockRes)
	if err != nil {
		return nil, fmt.Errorf("Fail to marshal:%s", err)
	}

	return &(blockRes.Block.Header.Time), nil
}
