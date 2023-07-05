package socials

import (
	"encoding/json"
	"fmt"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/singletons"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/src/service"
	"github.com/tranvannghia021/gocore/vars"
	"strconv"
	"time"
)

var shopify = "shopify"

type sShopify struct {
	http *service.SHttpRequest
}

var (
	scopeSp []string
	domain  string
)

type shopInfo struct {
	Shop struct {
		ID                                   int      `json:"id"`
		Name                                 string   `json:"name"`
		Email                                string   `json:"email"`
		Domain                               string   `json:"domain"`
		Province                             any      `json:"province"`
		Country                              string   `json:"country"`
		Address1                             string   `json:"address1"`
		Zip                                  any      `json:"zip"`
		City                                 any      `json:"city"`
		Source                               any      `json:"source"`
		Phone                                string   `json:"phone"`
		Latitude                             any      `json:"latitude"`
		Longitude                            any      `json:"longitude"`
		PrimaryLocale                        string   `json:"primary_locale"`
		Address2                             string   `json:"address2"`
		CreatedAt                            string   `json:"created_at"`
		UpdatedAt                            string   `json:"updated_at"`
		CountryCode                          string   `json:"country_code"`
		CountryName                          string   `json:"country_name"`
		Currency                             string   `json:"currency"`
		CustomerEmail                        string   `json:"customer_email"`
		Timezone                             string   `json:"timezone"`
		IanaTimezone                         string   `json:"iana_timezone"`
		ShopOwner                            string   `json:"shop_owner"`
		MoneyFormat                          string   `json:"money_format"`
		MoneyWithCurrencyFormat              string   `json:"money_with_currency_format"`
		WeightUnit                           string   `json:"weight_unit"`
		ProvinceCode                         any      `json:"province_code"`
		TaxesIncluded                        bool     `json:"taxes_included"`
		AutoConfigureTaxInclusivity          any      `json:"auto_configure_tax_inclusivity"`
		TaxShipping                          any      `json:"tax_shipping"`
		CountyTaxes                          bool     `json:"county_taxes"`
		PlanDisplayName                      string   `json:"plan_display_name"`
		PlanName                             string   `json:"plan_name"`
		HasDiscounts                         bool     `json:"has_discounts"`
		HasGiftCards                         bool     `json:"has_gift_cards"`
		MyshopifyDomain                      string   `json:"myshopify_domain"`
		GoogleAppsDomain                     any      `json:"google_apps_domain"`
		GoogleAppsLoginEnabled               any      `json:"google_apps_login_enabled"`
		MoneyInEmailsFormat                  string   `json:"money_in_emails_format"`
		MoneyWithCurrencyInEmailsFormat      string   `json:"money_with_currency_in_emails_format"`
		EligibleForPayments                  bool     `json:"eligible_for_payments"`
		RequiresExtraPaymentsAgreement       bool     `json:"requires_extra_payments_agreement"`
		PasswordEnabled                      bool     `json:"password_enabled"`
		HasStorefront                        bool     `json:"has_storefront"`
		Finances                             bool     `json:"finances"`
		PrimaryLocationID                    int64    `json:"primary_location_id"`
		CookieConsentLevel                   string   `json:"cookie_consent_level"`
		VisitorTrackingConsentPreference     string   `json:"visitor_tracking_consent_preference"`
		CheckoutAPISupported                 bool     `json:"checkout_api_supported"`
		MultiLocationEnabled                 bool     `json:"multi_location_enabled"`
		SetupRequired                        bool     `json:"setup_required"`
		PreLaunchEnabled                     bool     `json:"pre_launch_enabled"`
		EnabledPresentmentCurrencies         []string `json:"enabled_presentment_currencies"`
		TransactionalSmsDisabled             bool     `json:"transactional_sms_disabled"`
		MarketingSmsConsentEnabledAtCheckout bool     `json:"marketing_sms_consent_enabled_at_checkout"`
	} `json:"shop"`
}

func (s *sShopify) loadConfig() {
	coreConfig.Separator = ","
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"unauthenticated_read_product_listings",
	}, scopeSp...))
	domain = singletons.InstancePayload().Domain
	urlAuth = fmt.Sprintf("https://%s/admin/oauth/authorize", domain)
	s.http = service.NewHttpRequest()
}

func (s *sShopify) getToken(code string) vars.ResReq {
	body, _ := buildPayloadToken(code, true)
	s.http.Url = fmt.Sprintf("https://%s/admin/oauth/access_token", domain)
	s.http.FormData = body
	return s.http.PostFormDataRequest()
}

func (s *sShopify) profile(token string) repositories.Core {
	s.setParameter(domain)
	var headers = make(map[string]string)
	headers["X-Shopify-Access-Token"] = token
	s.http.Headers = headers
	s.http.Url = fmt.Sprintf("%s/shop.json", vars.EndPoint)
	result := s.http.GetRequest()
	if !result.Status {
		helpers.CheckNilErr(result.Error)
		return repositories.Core{}
	}
	var shopInfo shopInfo
	_ = json.Unmarshal(result.Data, &shopInfo)
	return repositories.Core{
		InternalId:    strconv.Itoa(shopInfo.Shop.ID),
		Platform:      shopify,
		Email:         shopInfo.Shop.Email,
		EmailVerifyAt: time.Now(),
		Password:      "",
		FirstName:     shopInfo.Shop.Name,
		Status:        true,
		Phone:         shopInfo.Shop.Phone,
		Address:       shopInfo.Shop.Address1,
		Domain:        shopInfo.Shop.Domain,
		RawDomain:     shopInfo.Shop.MyshopifyDomain,
	}
}

func (s *sShopify) setParameter(rawDomain string) {
	domain = rawDomain
	vars.EndPoint = fmt.Sprintf("https://%s/admin/api/%s", rawDomain, vars.Version)
}

func AddScopeShopify(scope []string) {
	scopeSp = scope
}
