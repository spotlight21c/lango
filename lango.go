package lango

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

var (
	localePath     string
	defaultLocale  string
	selectedLocale string
	translations   map[string]map[string]string
)

func Init(path, locale string) error {
	localePath = path

	translations = map[string]map[string]string{}

	files, err := ioutil.ReadDir(localePath)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errors.New("No language files found")
	}

	err = setDefaultLanguage(locale)
	if err != nil {
		return err
	}

	return nil
}

func setDefaultLanguage(locale string) error {
	defaultLocale = locale

	if err := loadLocaleFile(locale); err != nil {
		return err
	}

	return nil
}

func SetLocale(locale string) error {
	selectedLocale = locale

	if err := loadLocaleFile(locale); err != nil {
		return err
	}

	return nil
}

func loadLocaleFile(locale string) error {
	var jsonMap map[string]string

	body, err := ioutil.ReadFile(localePath + "/" + locale + ".json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &jsonMap)
	if err != nil {
		return err
	}

	translations[locale] = jsonMap

	return nil
}

func Get(key string, a ...interface{}) string {
	if v, ok := translations[selectedLocale][key]; ok {
		if len(a) > 0 {
			return fmt.Sprintf(v, a...)
		}

		return v
	}

	if v, ok := translations[defaultLocale][key]; ok {
		if len(a) > 0 {
			return fmt.Sprintf(v, a...)
		}

		return v
	}

	return key
}
