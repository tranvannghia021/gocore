package socials

import (
	"encoding/json"
	"fmt"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/src/service"
	"github.com/tranvannghia021/gocore/vars"
	"time"
)

var reddit = "reddit"

type sReddit struct {
	http *service.SHttpRequest
}

var scopeRd []string

type profileRd struct {
	IsEmployee           bool   `json:"is_employee"`
	SeenLayoutSwitch     bool   `json:"seen_layout_switch"`
	HasVisitedNewProfile bool   `json:"has_visited_new_profile"`
	PrefNoProfanity      bool   `json:"pref_no_profanity"`
	HasExternalAccount   bool   `json:"has_external_account"`
	PrefGeopopular       string `json:"pref_geopopular"`
	SeenRedesignModal    bool   `json:"seen_redesign_modal"`
	PrefShowTrending     bool   `json:"pref_show_trending"`
	Subreddit            struct {
		DefaultSet                 bool   `json:"default_set"`
		UserIsContributor          bool   `json:"user_is_contributor"`
		BannerImg                  string `json:"banner_img"`
		RestrictPosting            bool   `json:"restrict_posting"`
		UserIsBanned               bool   `json:"user_is_banned"`
		FreeFormReports            bool   `json:"free_form_reports"`
		CommunityIcon              any    `json:"community_icon"`
		ShowMedia                  bool   `json:"show_media"`
		IconColor                  string `json:"icon_color"`
		UserIsMuted                any    `json:"user_is_muted"`
		DisplayName                string `json:"display_name"`
		HeaderImg                  any    `json:"header_img"`
		Title                      string `json:"title"`
		Coins                      int    `json:"coins"`
		PreviousNames              []any  `json:"previous_names"`
		Over18                     bool   `json:"over_18"`
		IconSize                   []int  `json:"icon_size"`
		PrimaryColor               string `json:"primary_color"`
		IconImg                    string `json:"icon_img"`
		Description                string `json:"description"`
		AllowedMediaInComments     []any  `json:"allowed_media_in_comments"`
		SubmitLinkLabel            string `json:"submit_link_label"`
		HeaderSize                 any    `json:"header_size"`
		RestrictCommenting         bool   `json:"restrict_commenting"`
		Subscribers                int    `json:"subscribers"`
		SubmitTextLabel            string `json:"submit_text_label"`
		IsDefaultIcon              bool   `json:"is_default_icon"`
		LinkFlairPosition          string `json:"link_flair_position"`
		DisplayNamePrefixed        string `json:"display_name_prefixed"`
		KeyColor                   string `json:"key_color"`
		Name                       string `json:"name"`
		IsDefaultBanner            bool   `json:"is_default_banner"`
		URL                        string `json:"url"`
		Quarantine                 bool   `json:"quarantine"`
		BannerSize                 any    `json:"banner_size"`
		UserIsModerator            bool   `json:"user_is_moderator"`
		AcceptFollowers            bool   `json:"accept_followers"`
		PublicDescription          string `json:"public_description"`
		LinkFlairEnabled           bool   `json:"link_flair_enabled"`
		DisableContributorRequests bool   `json:"disable_contributor_requests"`
		SubredditType              string `json:"subreddit_type"`
		UserIsSubscriber           bool   `json:"user_is_subscriber"`
	} `json:"subreddit"`
	PrefShowPresence    bool   `json:"pref_show_presence"`
	SnoovatarImg        string `json:"snoovatar_img"`
	SnoovatarSize       []int  `json:"snoovatar_size"`
	GoldExpiration      any    `json:"gold_expiration"`
	HasGoldSubscription bool   `json:"has_gold_subscription"`
	IsSponsor           bool   `json:"is_sponsor"`
	NumFriends          int    `json:"num_friends"`
	Features            struct {
		ModmailHarassmentFilter   bool `json:"modmail_harassment_filter"`
		ModServiceMuteWrites      bool `json:"mod_service_mute_writes"`
		PromotedTrendBlanks       bool `json:"promoted_trend_blanks"`
		ShowAmpLink               bool `json:"show_amp_link"`
		Chat                      bool `json:"chat"`
		IsEmailPermissionRequired bool `json:"is_email_permission_required"`
		ModAwards                 bool `json:"mod_awards"`
		MwebXpromoRevampV3        struct {
			Owner        string `json:"owner"`
			Variant      string `json:"variant"`
			ExperimentID int    `json:"experiment_id"`
		} `json:"mweb_xpromo_revamp_v3"`
		MwebXpromoRevampV2 struct {
			Owner        string `json:"owner"`
			Variant      string `json:"variant"`
			ExperimentID int    `json:"experiment_id"`
		} `json:"mweb_xpromo_revamp_v2"`
		AwardsOnStreams                                bool `json:"awards_on_streams"`
		MwebXpromoModalListingClickDailyDismissibleIos bool `json:"mweb_xpromo_modal_listing_click_daily_dismissible_ios"`
		ChatSubreddit                                  bool `json:"chat_subreddit"`
		CookieConsentBanner                            bool `json:"cookie_consent_banner"`
		ModlogCopyrightRemoval                         bool `json:"modlog_copyright_removal"`
		DoNotTrack                                     bool `json:"do_not_track"`
		ImagesInComments                               bool `json:"images_in_comments"`
		ModServiceMuteReads                            bool `json:"mod_service_mute_reads"`
		ChatUserSettings                               bool `json:"chat_user_settings"`
		ResizedStylesImages                            bool `json:"resized_styles_images"`
		UsePrefAccountDeployment                       bool `json:"use_pref_account_deployment"`
		MwebXpromoInterstitialCommentsIos              bool `json:"mweb_xpromo_interstitial_comments_ios"`
		MwebLinkTab                                    struct {
			Owner        string `json:"owner"`
			Variant      string `json:"variant"`
			ExperimentID int    `json:"experiment_id"`
		} `json:"mweb_link_tab"`
		PremiumSubscriptionsTable                          bool `json:"premium_subscriptions_table"`
		MwebXpromoInterstitialCommentsAndroid              bool `json:"mweb_xpromo_interstitial_comments_android"`
		CrowdControlForPost                                bool `json:"crowd_control_for_post"`
		MwebXpromoModalListingClickDailyDismissibleAndroid bool `json:"mweb_xpromo_modal_listing_click_daily_dismissible_android"`
		ChatGroupRollout                                   bool `json:"chat_group_rollout"`
		MwebSharingClipboard                               struct {
			Owner        string `json:"owner"`
			Variant      string `json:"variant"`
			ExperimentID int    `json:"experiment_id"`
		} `json:"mweb_sharing_clipboard"`
		NoreferrerToNoopener bool `json:"noreferrer_to_noopener"`
		FeedAdLoad3          struct {
			Owner        string `json:"owner"`
			Variant      string `json:"variant"`
			ExperimentID int    `json:"experiment_id"`
		} `json:"feed_ad_load_3"`
		ExpensiveCoinsPackage bool `json:"expensive_coins_package"`
	} `json:"features"`
	CanEditName             bool     `json:"can_edit_name"`
	Verified                bool     `json:"verified"`
	PrefAutoplay            bool     `json:"pref_autoplay"`
	Coins                   int      `json:"coins"`
	HasPaypalSubscription   bool     `json:"has_paypal_subscription"`
	HasSubscribedToPremium  bool     `json:"has_subscribed_to_premium"`
	ID                      string   `json:"id"`
	HasStripeSubscription   bool     `json:"has_stripe_subscription"`
	OauthClientID           string   `json:"oauth_client_id"`
	CanCreateSubreddit      bool     `json:"can_create_subreddit"`
	Over18                  bool     `json:"over_18"`
	IsGold                  bool     `json:"is_gold"`
	IsMod                   bool     `json:"is_mod"`
	AwarderKarma            int      `json:"awarder_karma"`
	SuspensionExpirationUtc any      `json:"suspension_expiration_utc"`
	HasVerifiedEmail        bool     `json:"has_verified_email"`
	IsSuspended             bool     `json:"is_suspended"`
	PrefVideoAutoplay       bool     `json:"pref_video_autoplay"`
	HasAndroidSubscription  bool     `json:"has_android_subscription"`
	InRedesignBeta          bool     `json:"in_redesign_beta"`
	IconImg                 string   `json:"icon_img"`
	PrefNightmode           bool     `json:"pref_nightmode"`
	AwardeeKarma            int      `json:"awardee_karma"`
	HideFromRobots          bool     `json:"hide_from_robots"`
	PasswordSet             bool     `json:"password_set"`
	LinkKarma               int      `json:"link_karma"`
	ForcePasswordReset      bool     `json:"force_password_reset"`
	TotalKarma              int      `json:"total_karma"`
	SeenGiveAwardTooltip    bool     `json:"seen_give_award_tooltip"`
	InboxCount              int      `json:"inbox_count"`
	SeenPremiumAdblockModal bool     `json:"seen_premium_adblock_modal"`
	PrefTopKarmaSubreddits  bool     `json:"pref_top_karma_subreddits"`
	PrefShowSnoovatar       bool     `json:"pref_show_snoovatar"`
	Name                    string   `json:"name"`
	PrefClickgadget         int      `json:"pref_clickgadget"`
	Created                 float64  `json:"created"`
	GoldCreddits            int      `json:"gold_creddits"`
	CreatedUtc              float64  `json:"created_utc"`
	HasIosSubscription      bool     `json:"has_ios_subscription"`
	PrefShowTwitter         bool     `json:"pref_show_twitter"`
	InBeta                  bool     `json:"in_beta"`
	CommentKarma            int      `json:"comment_karma"`
	AcceptFollowers         bool     `json:"accept_followers"`
	HasSubscribed           bool     `json:"has_subscribed"`
	LinkedIdentities        []string `json:"linked_identities"`
	SeenSubredditChatFtux   bool     `json:"seen_subreddit_chat_ftux"`
}

func (s *sReddit) loadConfig() {
	coreConfig.Separator = " "
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"account",
		"identity",
	}, scopeRd...))
	urlAuth = "https://www.reddit.com/api/v1/authorize"
	parameters["duration"] = "permanent"
	s.http = service.NewHttpRequest()
}

func (s *sReddit) getToken(code string) vars.ResReq {
	body, _ := buildPayloadToken(code, true)
	s.http.Url = fmt.Sprintf("https://www.reddit.com/api/v1/access_token?%s", body.Encode())
	s.http.Headers = headerAuthBasic()
	return s.http.PostRequest()
}

func (s *sReddit) profile(token string) repositories.Core {
	s.http.Url = fmt.Sprintf("%s/api/%s/me", vars.EndPoint, vars.Version)
	result := s.http.SetAuth(token).GetRequest()
	if !result.Status {
		helpers.CheckNilErr(result.Error)
		return repositories.Core{}
	}
	var profile profileRd
	_ = json.Unmarshal(result.Data, &profile)
	return repositories.Core{
		InternalId:    profile.ID,
		Platform:      reddit,
		EmailVerifyAt: time.Now(),
		Password:      "",
		FirstName:     profile.Name,
		Avatar:        profile.SnoovatarImg,
		Status:        true,
	}
}

func AddScopeReddit(scope []string) {
	scopeRd = scope
}
