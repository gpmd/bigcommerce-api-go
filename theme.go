package bigcommerce

import (
	"encoding/json"
	"net/http"
	"time"
)

type ThemeConfig struct {
	UUID      string `json:"uuid"`
	StoreHash string `json:"store_hash"`
	ChannelID int64  `json:"channel_id"`
	Settings  struct {
		HideBreadcrumbs                                       bool     `json:"hide_breadcrumbs"`
		HidePageHeading                                       bool     `json:"hide_page_heading"`
		HideCategoryPageHeading                               bool     `json:"hide_category_page_heading"`
		HideBlogPageHeading                                   bool     `json:"hide_blog_page_heading"`
		HideContactUsPageHeading                              bool     `json:"hide_contact_us_page_heading"`
		HomepageNewProductsCount                              int      `json:"homepage_new_products_count"`
		HomepageFeaturedProductsCount                         int      `json:"homepage_featured_products_count"`
		HomepageTopProductsCount                              int      `json:"homepage_top_products_count"`
		HomepageShowCarousel                                  bool     `json:"homepage_show_carousel"`
		HomepageShowCarouselArrows                            bool     `json:"homepage_show_carousel_arrows"`
		HomepageShowCarouselPlayPauseButton                   bool     `json:"homepage_show_carousel_play_pause_button"`
		HomepageStretchCarouselImages                         bool     `json:"homepage_stretch_carousel_images"`
		HomepageNewProductsColumnCount                        int      `json:"homepage_new_products_column_count"`
		HomepageFeaturedProductsColumnCount                   int      `json:"homepage_featured_products_column_count"`
		HomepageTopProductsColumnCount                        int      `json:"homepage_top_products_column_count"`
		HomepageBlogPostsCount                                int      `json:"homepage_blog_posts_count"`
		ProductpageVideosCount                                int      `json:"productpage_videos_count"`
		ProductpageReviewsCount                               int      `json:"productpage_reviews_count"`
		ProductpageRelatedProductsCount                       int      `json:"productpage_related_products_count"`
		ProductpageSimilarByViewsCount                        int      `json:"productpage_similar_by_views_count"`
		CategorypageProductsPerPage                           int      `json:"categorypage_products_per_page"`
		ShopByPriceVisibility                                 bool     `json:"shop_by_price_visibility"`
		BrandpageProductsPerPage                              int      `json:"brandpage_products_per_page"`
		SearchpageProductsPerPage                             int      `json:"searchpage_products_per_page"`
		ShowProductQuickView                                  bool     `json:"show_product_quick_view"`
		ShowProductQuantityBox                                bool     `json:"show_product_quantity_box"`
		ShowPoweredBy                                         bool     `json:"show_powered_by"`
		ShopByBrandShowFooter                                 bool     `json:"shop_by_brand_show_footer"`
		ShowCopyrightFooter                                   bool     `json:"show_copyright_footer"`
		ShowAcceptAmex                                        bool     `json:"show_accept_amex"`
		ShowAcceptDiscover                                    bool     `json:"show_accept_discover"`
		ShowAcceptMastercard                                  bool     `json:"show_accept_mastercard"`
		ShowAcceptPaypal                                      bool     `json:"show_accept_paypal"`
		ShowAcceptVisa                                        bool     `json:"show_accept_visa"`
		ShowAcceptAmazonpay                                   bool     `json:"show_accept_amazonpay"`
		ShowAcceptGooglepay                                   bool     `json:"show_accept_googlepay"`
		ShowAcceptKlarna                                      bool     `json:"show_accept_klarna"`
		ShowProductDetailsTabs                                bool     `json:"show_product_details_tabs"`
		ShowProductReviews                                    bool     `json:"show_product_reviews"`
		ShowCustomFieldsTabs                                  bool     `json:"show_custom_fields_tabs"`
		ShowProductWeight                                     bool     `json:"show_product_weight"`
		ShowProductDimensions                                 bool     `json:"show_product_dimensions"`
		ShowProductSwatchNames                                bool     `json:"show_product_swatch_names"`
		ProductListDisplayMode                                string   `json:"product_list_display_mode"`
		LogoPosition                                          string   `json:"logo-position"`
		LogoSize                                              string   `json:"logo_size"`
		LogoFontSize                                          int      `json:"logo_fontSize"`
		BrandSize                                             string   `json:"brand_size"`
		GallerySize                                           string   `json:"gallery_size"`
		ProductgallerySize                                    string   `json:"productgallery_size"`
		ProductSize                                           string   `json:"product_size"`
		ProductviewThumbSize                                  string   `json:"productview_thumb_size"`
		ProductthumbSize                                      string   `json:"productthumb_size"`
		ThumbSize                                             string   `json:"thumb_size"`
		ZoomSize                                              string   `json:"zoom_size"`
		BlogSize                                              string   `json:"blog_size"`
		DefaultImageBrand                                     string   `json:"default_image_brand"`
		DefaultImageProduct                                   string   `json:"default_image_product"`
		DefaultImageGiftCertificate                           string   `json:"default_image_gift_certificate"`
		BodyFont                                              string   `json:"body-font"`
		HeadingsFont                                          string   `json:"headings-font"`
		FontSizeRoot                                          int      `json:"fontSize-root"`
		FontSizeH1                                            int      `json:"fontSize-h1"`
		FontSizeH2                                            int      `json:"fontSize-h2"`
		FontSizeH3                                            int      `json:"fontSize-h3"`
		FontSizeH4                                            int      `json:"fontSize-h4"`
		FontSizeH5                                            int      `json:"fontSize-h5"`
		FontSizeH6                                            int      `json:"fontSize-h6"`
		ApplePayButton                                        string   `json:"applePay-button"`
		ColorTextBase                                         string   `json:"color-textBase"`
		ColorTextBaseHover                                    string   `json:"color-textBase--hover"`
		ColorTextBaseActive                                   string   `json:"color-textBase--active"`
		ColorTextSecondary                                    string   `json:"color-textSecondary"`
		ColorTextSecondaryHover                               string   `json:"color-textSecondary--hover"`
		ColorTextSecondaryActive                              string   `json:"color-textSecondary--active"`
		ColorTextLink                                         string   `json:"color-textLink"`
		ColorTextLinkHover                                    string   `json:"color-textLink--hover"`
		ColorTextLinkActive                                   string   `json:"color-textLink--active"`
		ColorTextHeading                                      string   `json:"color-textHeading"`
		ColorPrimary                                          string   `json:"color-primary"`
		ColorPrimaryDark                                      string   `json:"color-primaryDark"`
		ColorPrimaryDarker                                    string   `json:"color-primaryDarker"`
		ColorPrimaryLight                                     string   `json:"color-primaryLight"`
		ColorSecondary                                        string   `json:"color-secondary"`
		ColorSecondaryDark                                    string   `json:"color-secondaryDark"`
		ColorSecondaryDarker                                  string   `json:"color-secondaryDarker"`
		ColorError                                            string   `json:"color-error"`
		ColorErrorLight                                       string   `json:"color-errorLight"`
		ColorInfo                                             string   `json:"color-info"`
		ColorInfoLight                                        string   `json:"color-infoLight"`
		ColorSuccess                                          string   `json:"color-success"`
		ColorSuccessLight                                     string   `json:"color-successLight"`
		ColorWarning                                          string   `json:"color-warning"`
		ColorWarningLight                                     string   `json:"color-warningLight"`
		ColorBlack                                            string   `json:"color-black"`
		ColorWhite                                            string   `json:"color-white"`
		ColorWhitesBase                                       string   `json:"color-whitesBase"`
		ColorGrey                                             string   `json:"color-grey"`
		ColorGreyDarkest                                      string   `json:"color-greyDarkest"`
		ColorGreyDarker                                       string   `json:"color-greyDarker"`
		ColorGreyDark                                         string   `json:"color-greyDark"`
		ColorGreyMedium                                       string   `json:"color-greyMedium"`
		ColorGreyLight                                        string   `json:"color-greyLight"`
		ColorGreyLighter                                      string   `json:"color-greyLighter"`
		ColorGreyLightest                                     string   `json:"color-greyLightest"`
		BannerDeafaultBackgroundColor                         string   `json:"banner--deafault-backgroundColor"`
		ButtonDefaultColor                                    string   `json:"button--default-color"`
		ButtonDefaultColorHover                               string   `json:"button--default-colorHover"`
		ButtonDefaultColorActive                              string   `json:"button--default-colorActive"`
		ButtonDefaultBorderColor                              string   `json:"button--default-borderColor"`
		ButtonDefaultBorderColorHover                         string   `json:"button--default-borderColorHover"`
		ButtonDefaultBorderColorActive                        string   `json:"button--default-borderColorActive"`
		ButtonPrimaryColor                                    string   `json:"button--primary-color"`
		ButtonPrimaryColorHover                               string   `json:"button--primary-colorHover"`
		ButtonPrimaryColorActive                              string   `json:"button--primary-colorActive"`
		ButtonPrimaryBackgroundColor                          string   `json:"button--primary-backgroundColor"`
		ButtonPrimaryBackgroundColorHover                     string   `json:"button--primary-backgroundColorHover"`
		ButtonPrimaryBackgroundColorActive                    string   `json:"button--primary-backgroundColorActive"`
		ButtonDisabledColor                                   string   `json:"button--disabled-color"`
		ButtonDisabledBackgroundColor                         string   `json:"button--disabled-backgroundColor"`
		ButtonDisabledBorderColor                             string   `json:"button--disabled-borderColor"`
		IconColor                                             string   `json:"icon-color"`
		IconColorHover                                        string   `json:"icon-color-hover"`
		ButtonIconSvgColor                                    string   `json:"button--icon-svg-color"`
		IconRatingEmpty                                       string   `json:"icon-ratingEmpty"`
		IconRatingFull                                        string   `json:"icon-ratingFull"`
		CarouselBgColor                                       string   `json:"carousel-bgColor"`
		CarouselTitleColor                                    string   `json:"carousel-title-color"`
		CarouselDescriptionColor                              string   `json:"carousel-description-color"`
		CarouselDotColor                                      string   `json:"carousel-dot-color"`
		CarouselDotColorActive                                string   `json:"carousel-dot-color-active"`
		CarouselDotBgColor                                    string   `json:"carousel-dot-bgColor"`
		CarouselArrowColor                                    string   `json:"carousel-arrow-color"`
		CarouselArrowColorHover                               string   `json:"carousel-arrow-color--hover"`
		CarouselArrowBgColor                                  string   `json:"carousel-arrow-bgColor"`
		CarouselArrowBorderColor                              string   `json:"carousel-arrow-borderColor"`
		CarouselPlayPauseButtonTextColor                      string   `json:"carousel-play-pause-button-textColor"`
		CarouselPlayPauseButtonTextColorHover                 string   `json:"carousel-play-pause-button-textColor--hover"`
		CarouselPlayPauseButtonBgColor                        string   `json:"carousel-play-pause-button-bgColor"`
		CarouselPlayPauseButtonBorderColor                    string   `json:"carousel-play-pause-button-borderColor"`
		CardTitleColor                                        string   `json:"card-title-color"`
		CardTitleColorHover                                   string   `json:"card-title-color-hover"`
		CardFigcaptionButtonBackground                        string   `json:"card-figcaption-button-background"`
		CardFigcaptionButtonColor                             string   `json:"card-figcaption-button-color"`
		CardAlternateBackgroundColor                          string   `json:"card--alternate-backgroundColor"`
		CardAlternateBorderColor                              string   `json:"card--alternate-borderColor"`
		CardAlternateColorHover                               string   `json:"card--alternate-color--hover"`
		FormLabelFontColor                                    string   `json:"form-label-font-color"`
		InputFontColor                                        string   `json:"input-font-color"`
		InputBorderColor                                      string   `json:"input-border-color"`
		InputBorderColorActive                                string   `json:"input-border-color-active"`
		InputBgColor                                          string   `json:"input-bg-color"`
		InputDisabledBg                                       string   `json:"input-disabled-bg"`
		SelectBgColor                                         string   `json:"select-bg-color"`
		SelectArrowColor                                      string   `json:"select-arrow-color"`
		CheckRadioColor                                       string   `json:"checkRadio-color"`
		CheckRadioBackgroundColor                             string   `json:"checkRadio-backgroundColor"`
		CheckRadioBorderColor                                 string   `json:"checkRadio-borderColor"`
		AlertBackgroundColor                                  string   `json:"alert-backgroundColor"`
		AlertColor                                            string   `json:"alert-color"`
		AlertColorAlt                                         string   `json:"alert-color-alt"`
		StoreNameColor                                        string   `json:"storeName-color"`
		BodyBg                                                string   `json:"body-bg"`
		HeaderBackgroundColor                                 string   `json:"header-backgroundColor"`
		FooterBackgroundColor                                 string   `json:"footer-backgroundColor"`
		NavUserColor                                          string   `json:"navUser-color"`
		NavUserColorHover                                     string   `json:"navUser-color-hover"`
		NavUserDropdownBackgroundColor                        string   `json:"navUser-dropdown-backgroundColor"`
		NavUserDropdownBorderColor                            string   `json:"navUser-dropdown-borderColor"`
		NavUserIndicatorBackgroundColor                       string   `json:"navUser-indicator-backgroundColor"`
		NavPagesColor                                         string   `json:"navPages-color"`
		NavPagesColorHover                                    string   `json:"navPages-color-hover"`
		NavPagesSubMenuBackgroundColor                        string   `json:"navPages-subMenu-backgroundColor"`
		NavPagesSubMenuSeparatorColor                         string   `json:"navPages-subMenu-separatorColor"`
		DropdownQuickSearchBackgroundColor                    string   `json:"dropdown--quickSearch-backgroundColor"`
		DropdownWishListBackgroundColor                       string   `json:"dropdown--wishList-backgroundColor"`
		BlockquoteCiteFontColor                               string   `json:"blockquote-cite-font-color"`
		ContainerBorderGlobalColorBase                        string   `json:"container-border-global-color-base"`
		ContainerFillBase                                     string   `json:"container-fill-base"`
		ContainerFillDark                                     string   `json:"container-fill-dark"`
		LabelBackgroundColor                                  string   `json:"label-backgroundColor"`
		LabelColor                                            string   `json:"label-color"`
		OverlayBackgroundColor                                string   `json:"overlay-backgroundColor"`
		LoadingOverlayBackgroundColor                         string   `json:"loadingOverlay-backgroundColor"`
		PaceProgressBackgroundColor                           string   `json:"pace-progress-backgroundColor"`
		SpinnerBorderColorDark                                string   `json:"spinner-borderColor-dark"`
		SpinnerBorderColorLight                               string   `json:"spinner-borderColor-light"`
		HideContentNavigation                                 bool     `json:"hide_content_navigation"`
		OptimizedCheckoutHeaderBackgroundColor                string   `json:"optimizedCheckout-header-backgroundColor"`
		OptimizedCheckoutShowBackgroundImage                  bool     `json:"optimizedCheckout-show-backgroundImage"`
		OptimizedCheckoutBackgroundImage                      string   `json:"optimizedCheckout-backgroundImage"`
		OptimizedCheckoutBackgroundImageSize                  string   `json:"optimizedCheckout-backgroundImage-size"`
		OptimizedCheckoutShowLogo                             string   `json:"optimizedCheckout-show-logo"`
		OptimizedCheckoutLogo                                 string   `json:"optimizedCheckout-logo"`
		OptimizedCheckoutLogoSize                             string   `json:"optimizedCheckout-logo-size"`
		OptimizedCheckoutLogoPosition                         string   `json:"optimizedCheckout-logo-position"`
		OptimizedCheckoutHeadingPrimaryColor                  string   `json:"optimizedCheckout-headingPrimary-color"`
		OptimizedCheckoutHeadingPrimaryFont                   string   `json:"optimizedCheckout-headingPrimary-font"`
		OptimizedCheckoutHeadingSecondaryColor                string   `json:"optimizedCheckout-headingSecondary-color"`
		OptimizedCheckoutHeadingSecondaryFont                 string   `json:"optimizedCheckout-headingSecondary-font"`
		OptimizedCheckoutBodyBackgroundColor                  string   `json:"optimizedCheckout-body-backgroundColor"`
		OptimizedCheckoutColorFocus                           string   `json:"optimizedCheckout-colorFocus"`
		OptimizedCheckoutContentPrimaryColor                  string   `json:"optimizedCheckout-contentPrimary-color"`
		OptimizedCheckoutContentPrimaryFont                   string   `json:"optimizedCheckout-contentPrimary-font"`
		OptimizedCheckoutContentSecondaryColor                string   `json:"optimizedCheckout-contentSecondary-color"`
		OptimizedCheckoutContentSecondaryFont                 string   `json:"optimizedCheckout-contentSecondary-font"`
		OptimizedCheckoutButtonPrimaryFont                    string   `json:"optimizedCheckout-buttonPrimary-font"`
		OptimizedCheckoutButtonPrimaryColor                   string   `json:"optimizedCheckout-buttonPrimary-color"`
		OptimizedCheckoutButtonPrimaryColorHover              string   `json:"optimizedCheckout-buttonPrimary-colorHover"`
		OptimizedCheckoutButtonPrimaryColorActive             string   `json:"optimizedCheckout-buttonPrimary-colorActive"`
		OptimizedCheckoutButtonPrimaryBackgroundColor         string   `json:"optimizedCheckout-buttonPrimary-backgroundColor"`
		OptimizedCheckoutButtonPrimaryBackgroundColorHover    string   `json:"optimizedCheckout-buttonPrimary-backgroundColorHover"`
		OptimizedCheckoutButtonPrimaryBackgroundColorActive   string   `json:"optimizedCheckout-buttonPrimary-backgroundColorActive"`
		OptimizedCheckoutButtonPrimaryBorderColor             string   `json:"optimizedCheckout-buttonPrimary-borderColor"`
		OptimizedCheckoutButtonPrimaryBorderColorHover        string   `json:"optimizedCheckout-buttonPrimary-borderColorHover"`
		OptimizedCheckoutButtonPrimaryBorderColorActive       string   `json:"optimizedCheckout-buttonPrimary-borderColorActive"`
		OptimizedCheckoutButtonPrimaryBorderColorDisabled     string   `json:"optimizedCheckout-buttonPrimary-borderColorDisabled"`
		OptimizedCheckoutButtonPrimaryBackgroundColorDisabled string   `json:"optimizedCheckout-buttonPrimary-backgroundColorDisabled"`
		OptimizedCheckoutButtonPrimaryColorDisabled           string   `json:"optimizedCheckout-buttonPrimary-colorDisabled"`
		OptimizedCheckoutFormChecklistBackgroundColor         string   `json:"optimizedCheckout-formChecklist-backgroundColor"`
		OptimizedCheckoutFormChecklistColor                   string   `json:"optimizedCheckout-formChecklist-color"`
		OptimizedCheckoutFormChecklistBorderColor             string   `json:"optimizedCheckout-formChecklist-borderColor"`
		OptimizedCheckoutFormChecklistBackgroundColorSelected string   `json:"optimizedCheckout-formChecklist-backgroundColorSelected"`
		OptimizedCheckoutButtonSecondaryFont                  string   `json:"optimizedCheckout-buttonSecondary-font"`
		OptimizedCheckoutButtonSecondaryColor                 string   `json:"optimizedCheckout-buttonSecondary-color"`
		OptimizedCheckoutButtonSecondaryColorHover            string   `json:"optimizedCheckout-buttonSecondary-colorHover"`
		OptimizedCheckoutButtonSecondaryColorActive           string   `json:"optimizedCheckout-buttonSecondary-colorActive"`
		OptimizedCheckoutButtonSecondaryBackgroundColor       string   `json:"optimizedCheckout-buttonSecondary-backgroundColor"`
		OptimizedCheckoutButtonSecondaryBorderColor           string   `json:"optimizedCheckout-buttonSecondary-borderColor"`
		OptimizedCheckoutButtonSecondaryBackgroundColorHover  string   `json:"optimizedCheckout-buttonSecondary-backgroundColorHover"`
		OptimizedCheckoutButtonSecondaryBackgroundColorActive string   `json:"optimizedCheckout-buttonSecondary-backgroundColorActive"`
		OptimizedCheckoutButtonSecondaryBorderColorHover      string   `json:"optimizedCheckout-buttonSecondary-borderColorHover"`
		OptimizedCheckoutButtonSecondaryBorderColorActive     string   `json:"optimizedCheckout-buttonSecondary-borderColorActive"`
		OptimizedCheckoutLinkColor                            string   `json:"optimizedCheckout-link-color"`
		OptimizedCheckoutLinkFont                             string   `json:"optimizedCheckout-link-font"`
		OptimizedCheckoutDiscountBannerBackgroundColor        string   `json:"optimizedCheckout-discountBanner-backgroundColor"`
		OptimizedCheckoutDiscountBannerTextColor              string   `json:"optimizedCheckout-discountBanner-textColor"`
		OptimizedCheckoutDiscountBannerIconColor              string   `json:"optimizedCheckout-discountBanner-iconColor"`
		OptimizedCheckoutOrderSummaryBackgroundColor          string   `json:"optimizedCheckout-orderSummary-backgroundColor"`
		OptimizedCheckoutOrderSummaryBorderColor              string   `json:"optimizedCheckout-orderSummary-borderColor"`
		OptimizedCheckoutStepBackgroundColor                  string   `json:"optimizedCheckout-step-backgroundColor"`
		OptimizedCheckoutStepTextColor                        string   `json:"optimizedCheckout-step-textColor"`
		OptimizedCheckoutFormTextColor                        string   `json:"optimizedCheckout-form-textColor"`
		OptimizedCheckoutFormFieldBorderColor                 string   `json:"optimizedCheckout-formField-borderColor"`
		OptimizedCheckoutFormFieldTextColor                   string   `json:"optimizedCheckout-formField-textColor"`
		OptimizedCheckoutFormFieldShadowColor                 string   `json:"optimizedCheckout-formField-shadowColor"`
		OptimizedCheckoutFormFieldPlaceholderColor            string   `json:"optimizedCheckout-formField-placeholderColor"`
		OptimizedCheckoutFormFieldBackgroundColor             string   `json:"optimizedCheckout-formField-backgroundColor"`
		OptimizedCheckoutFormFieldErrorColor                  string   `json:"optimizedCheckout-formField-errorColor"`
		OptimizedCheckoutFormFieldInputControlColor           string   `json:"optimizedCheckout-formField-inputControlColor"`
		OptimizedCheckoutStepBorderColor                      string   `json:"optimizedCheckout-step-borderColor"`
		OptimizedCheckoutHeaderBorderColor                    string   `json:"optimizedCheckout-header-borderColor"`
		OptimizedCheckoutHeaderTextColor                      string   `json:"optimizedCheckout-header-textColor"`
		OptimizedCheckoutLoadingToasterBackgroundColor        string   `json:"optimizedCheckout-loadingToaster-backgroundColor"`
		OptimizedCheckoutLoadingToasterTextColor              string   `json:"optimizedCheckout-loadingToaster-textColor"`
		OptimizedCheckoutLinkHoverColor                       string   `json:"optimizedCheckout-link-hoverColor"`
		ProductSaleBadges                                     string   `json:"product_sale_badges"`
		ColorBadgeProductSaleBadges                           string   `json:"color_badge_product_sale_badges"`
		ColorTextProductSaleBadges                            string   `json:"color_text_product_sale_badges"`
		ColorHoverProductSaleBadges                           string   `json:"color_hover_product_sale_badges"`
		ProductSoldOutBadges                                  string   `json:"product_sold_out_badges"`
		ColorBadgeProductSoldOutBadges                        string   `json:"color_badge_product_sold_out_badges"`
		ColorTextProductSoldOutBadges                         string   `json:"color_text_product_sold_out_badges"`
		ColorHoverProductSoldOutBadges                        string   `json:"color_hover_product_sold_out_badges"`
		FocusTooltipTextColor                                 string   `json:"focusTooltip-textColor"`
		FocusTooltipBackgroundColor                           string   `json:"focusTooltip-backgroundColor"`
		RestrictToLogin                                       bool     `json:"restrict_to_login"`
		SwatchOptionSize                                      string   `json:"swatch_option_size"`
		SocialIconPlacementTop                                bool     `json:"social_icon_placement_top"`
		SocialIconPlacementBottom                             string   `json:"social_icon_placement_bottom"`
		NavigationDesign                                      string   `json:"navigation_design"`
		PriceRanges                                           bool     `json:"price_ranges"`
		PdpPriceLabel                                         string   `json:"pdp-price-label"`
		PdpSaleBadgeLabel                                     string   `json:"pdp_sale_badge_label"`
		PdpSoldOutLabel                                       string   `json:"pdp_sold_out_label"`
		PdpSalePriceLabel                                     string   `json:"pdp-sale-price-label"`
		PdpNonSalePriceLabel                                  string   `json:"pdp-non-sale-price-label"`
		PdpRetailPriceLabel                                   string   `json:"pdp-retail-price-label"`
		PdpCustomFieldsTabLabel                               string   `json:"pdp-custom-fields-tab-label"`
		PaymentbuttonsPaypalLayout                            string   `json:"paymentbuttons-paypal-layout"`
		PaymentbuttonsPaypalColor                             string   `json:"paymentbuttons-paypal-color"`
		PaymentbuttonsPaypalShape                             string   `json:"paymentbuttons-paypal-shape"`
		PaymentbuttonsPaypalLabel                             string   `json:"paymentbuttons-paypal-label"`
		PaymentbannersHomepageColor                           string   `json:"paymentbanners-homepage-color"`
		PaymentbannersHomepageRatio                           string   `json:"paymentbanners-homepage-ratio"`
		PaymentbannersCartpageTextColor                       string   `json:"paymentbanners-cartpage-text-color"`
		PaymentbannersCartpageLogoPosition                    string   `json:"paymentbanners-cartpage-logo-position"`
		PaymentbannersCartpageLogoType                        string   `json:"paymentbanners-cartpage-logo-type"`
		PaymentbannersProddetailspageColor                    string   `json:"paymentbanners-proddetailspage-color"`
		PaymentbannersProddetailspageRatio                    string   `json:"paymentbanners-proddetailspage-ratio"`
		PaymentbuttonsContainer                               string   `json:"paymentbuttons-container"`
		SupportedCardTypeIcons                                []string `json:"supported_card_type_icons"`
		SupportedPaymentMethods                               []string `json:"supported_payment_methods"`
		LazyloadMode                                          string   `json:"lazyload_mode"`
		CheckoutPaymentbuttonsPaypalColor                     string   `json:"checkout-paymentbuttons-paypal-color"`
		CheckoutPaymentbuttonsPaypalShape                     string   `json:"checkout-paymentbuttons-paypal-shape"`
		CheckoutPaymentbuttonsPaypalSize                      string   `json:"checkout-paymentbuttons-paypal-size"`
		CheckoutPaymentbuttonsPaypalLabel                     string   `json:"checkout-paymentbuttons-paypal-label"`
	} `json:"settings"`
	ThemeUUID     string    `json:"theme_uuid"`
	VersionUUID   string    `json:"version_uuid"`
	VariationUUID string    `json:"variation_uuid"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Theme struct {
	UUID       string `json:"uuid"`
	Variations []struct {
		UUID        string `json:"uuid"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ExternalID  string `json:"external_id"`
	} `json:"variations"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"is_private"`
	IsActive  bool   `json:"is_active"`
}

// GetActiveThemeConfig returns the active theme config (not handling variations yet)
func (bc *Client) GetActiveThemeConfig() (*ThemeConfig, error) {
	var themeConfig ThemeConfig
	themes, err := bc.GetThemes()
	if err != nil {
		return nil, err
	}
	for _, theme := range themes {
		if theme.IsActive {
			return bc.GetThemeConfig(theme.UUID)
		}
	}
	return &themeConfig, nil
}

// GetThemes returns a list of all store themes
func (bc *Client) GetThemes() ([]Theme, error) {
	req := bc.getAPIRequest(http.MethodGet, "/v3/themes", nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := processBody(res)
	if err != nil {
		return nil, err
	}

	var ret struct {
		Data []Theme `json:"data"`
	}
	err = json.Unmarshal(body, &ret)
	return ret.Data, err

}

// GetThemeConfig returns the configuration for a specific theme by theme UUID
func (bc *Client) GetThemeConfig(uuid string) (*ThemeConfig, error) {
	req := bc.getAPIRequest(http.MethodGet, "/v3/themes/"+uuid+"/configurations", nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := processBody(res)
	if err != nil {
		return nil, err
	}

	var ret struct {
		Data []ThemeConfig `json:"data"`
	}
	err = json.Unmarshal(body, &ret)
	return &ret.Data[0], err
}
