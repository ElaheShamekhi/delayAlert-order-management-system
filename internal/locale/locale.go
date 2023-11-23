package locale

import (
	"delayAlert-order-management-system/internal/config"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
	"path"
)

type Localizer struct {
	locales map[language.Tag]*i18n.Localizer
}

var l *Localizer

func Localize(msgID string, lang language.Tag) string {
	if loc, ok := l.locales[lang]; ok {
		localize, err := loc.Localize(&i18n.LocalizeConfig{MessageID: msgID})
		if err != nil || localize == "" {
			return msgID
		}
		return localize
	}
	return msgID
}

func Init() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err := bundle.LoadMessageFile(path.Join(config.LocalePath(), "validations.en.toml"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load validations.en.toml file")
	}
	_, err = bundle.LoadMessageFile(path.Join(config.LocalePath(), "validations.fa.toml"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load validations.fa.toml file")
	}
	_, err = bundle.LoadMessageFile(path.Join(config.LocalePath(), "errors.fa.toml"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load errors.fa.toml file")
	}
	l = &Localizer{locales: make(map[language.Tag]*i18n.Localizer)}
	l.locales[language.Persian] = i18n.NewLocalizer(bundle, "fa")
	l.locales[language.English] = i18n.NewLocalizer(bundle, "en")
}
