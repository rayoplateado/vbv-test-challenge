package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"factorialfrontend/log"
	"io"
	"net/http"
	"os"
	"strconv"
)

type IApi[T any] interface {
	Get(resource string) (T, error)
	List(resource string) ([]T, error)
	Post(resource string, payload any) (T, error)
}

type Api[T any] struct {
}

func (a Api[T]) Get(resource string) (T, error) {
	var err error
	var t T

	resp, err := doRequest("GET", resource, nil)

	if err != nil {
		log.Error("error making request: " + err.Error())
		return t, err
	}

	defer resp.Body.Close()

	t, err = validateUnmarshal[T](resp)

	if err != nil {
		log.Error("error unmarshaling: " + err.Error())
		return t, err
	}

	return t, err
}

func (a Api[T]) List(resource string) ([]T, error) {
	var err error
	var list []T

	resp, err := doRequest("GET", resource, nil)

	if err != nil {
		log.Error("error making request: " + err.Error())
		return list, err
	}

	defer resp.Body.Close()

	list, err = validateUnmarshal[[]T](resp)

	if err != nil {
		log.Error("error unmarshaling: " + err.Error())
		return list, err
	}

	return list, err
}

func (a Api[T]) Post(resource string, payload any) (T, error) {
	var err error
	var t T

	resp, err := doRequest("POST", resource, payload)

	if err != nil {
		log.Error("error making request: " + err.Error())
		return t, err
	}

	defer resp.Body.Close()

	t, err = validateUnmarshal[T](resp)

	if err != nil {
		log.Error("error unmarshaling: " + err.Error())
		return t, err
	}

	return t, err
}

func doRequest(method, resource string, payload any) (*http.Response, error) {
	client := &http.Client{}
	var jsonData []byte
	var err error

	if payload != nil {
		jsonData, err = json.Marshal(payload)
		if err != nil {
			return nil, err
		}
	}

	log.Info("making " + method + " request to " + os.Getenv("APIADDRESS") + resource)
	req, err := http.NewRequest(method, os.Getenv("APIADDRESS")+resource, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Error("error with the request done: " + err.Error())
		return nil, err
	}

	return resp, nil
}

func validateUnmarshal[T any](resp *http.Response) (T, error) {
	var result T

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Error("Unexpected status code when unmarshaling response: " + strconv.Itoa(resp.StatusCode))
		return result, errors.New("Unexpected status code: " + strconv.Itoa(resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("error reading bytes from response: " + err.Error())
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error("error unmarshaling bytes to json: " + err.Error())
		return result, err
	}

	return result, nil
}
