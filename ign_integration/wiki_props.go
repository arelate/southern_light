package ign_integration

import (
	"fmt"
	"github.com/boggydigital/match_node"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
	"time"
)

const (
	wikisLinkPfx = "/wikis"
)

type Breadcrumb struct {
	Typename string `json:"__typename"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Type     string `json:"type"`
	Paywall  bool   `json:"paywall"`
	Metadata struct {
		Typename string `json:"__typename"`
		Names    struct {
			Typename string `json:"__typename"`
			Name     string `json:"name"`
		} `json:"names"`
	} `json:"metadata"`
	PrimaryImage struct {
		Typename string `json:"__typename"`
		Url      string `json:"url"`
	} `json:"primaryImage"`
	Genres []struct {
		Typename string `json:"__typename"`
		Name     string `json:"name"`
		Slug     string `json:"slug"`
	} `json:"genres"`
	ObjectRegions []struct {
		Typename string `json:"__typename"`
		Id       int    `json:"id"`
		Releases []struct {
			Typename           string `json:"__typename"`
			Id                 string `json:"id"`
			PlatformAttributes []struct {
				Typename string `json:"__typename"`
				Id       int    `json:"id"`
				Name     string `json:"name"`
				Slug     string `json:"slug"`
			} `json:"platformAttributes"`
		} `json:"releases"`
	} `json:"objectRegions"`
	Franchises []interface{} `json:"franchises"`
}

type HTMLEntity struct {
	Typename string `json:"__typename"`
	Name     string `json:"name"`
	Values   struct {
		Html   string `json:"html"`
		Tag    string `json:"tag"`
		Header bool   `json:"header"`
		Block  bool   `json:"block"`
	} `json:"values"`
}

type Views struct {
	TOC           string `json:"TOC"`
	Search        string `json:"Search"`
	Settings      string `json:"Settings"`
	Maps          string `json:"Maps"`
	MapContent    string `json:"MapContent"`
	Notes         string `json:"Notes"`
	CustomMarkers string `json:"CustomMarkers"`
	Regions       string `json:"Regions"`
	Groups        string `json:"Groups"`
	GroupMarkers  string `json:"GroupMarkers"`
	AddMarker     string `json:"AddMarker"`
	EditMarker    string `json:"EditMarker"`
	Checklists    string `json:"Checklists"`
	Heatmaps      string `json:"Heatmaps"`
}

type Contributor struct {
	Typename    string  `json:"__typename"`
	Name        *string `json:"name"`
	Nickname    string  `json:"nickname"`
	DisplayName string  `json:"displayName"`
}

type SiblingPage struct {
	Typename string `json:"__typename"`
	Label    string `json:"label"`
	Url      string `json:"url"`
}

type WikiProps struct {
	Props struct {
		PageProps struct {
			Page struct {
				Canonical           string        `json:"canonical"`
				Slug                string        `json:"slug"`
				Typename            string        `json:"__typename"`
				Id                  string        `json:"id"`
				Name                string        `json:"name"`
				Tags                interface{}   `json:"tags"`
				Locked              bool          `json:"locked"`
				Contributors        int           `json:"contributors"`
				Maps                []interface{} `json:"maps"`
				TopContributors     []Contributor `json:"topContributors"`
				Breadcrumbs         []Breadcrumb  `json:"breadcrumbs"`
				IsCommunity         string        `json:"isCommunity"`
				Resource            string        `json:"resource"`
				PageType            string        `json:"pageType"`
				Chapter             string        `json:"chapter"`
				ChapterSlug         string        `json:"chapterSlug"`
				Description         string        `json:"description"`
				Title               string        `json:"title"`
				AssetId             string        `json:"assetId"`
				GenreSlugs          []string      `json:"genreSlugs"`
				Platforms           []string      `json:"platforms"`
				PlatformSlugs       []string      `json:"platformSlugs"`
				ObjectId            string        `json:"objectId"`
				ObjectSlug          string        `json:"objectSlug"`
				UpdatedAt           time.Time     `json:"updatedAt"`
				PublishDate         time.Time     `json:"publishDate"`
				AdditionalDataLayer struct {
					Content struct {
						PageName  string   `json:"pageName"`
						ContentId string   `json:"contentId"`
						ObjectIds []string `json:"objectIds"`
					} `json:"content"`
				} `json:"additionalDataLayer"`
				Page struct {
					Typename       string        `json:"__typename"`
					Id             string        `json:"id"`
					Title          string        `json:"title"`
					UpdatedAt      time.Time     `json:"updatedAt"`
					PublishDate    time.Time     `json:"publishDate"`
					Locked         bool          `json:"locked"`
					RevisionId     int           `json:"revisionId"`
					Thin           bool          `json:"thin"`
					HasVideo       bool          `json:"hasVideo"`
					ImageForTitle  string        `json:"imageForTitle"`
					MapMarker      interface{}   `json:"mapMarker"`
					HtmlEntities   []HTMLEntity  `json:"htmlEntities"`
					NextPage       SiblingPage   `json:"nextPage"`
					PrevPage       SiblingPage   `json:"prevPage"`
					ChecklistTasks []interface{} `json:"checklistTasks"`
				} `json:"page"`
				PrimaryObject struct {
					Typename string `json:"__typename"`
					Id       string `json:"id"`
					Slug     string `json:"slug"`
					Type     string `json:"type"`
					Paywall  bool   `json:"paywall"`
					Metadata struct {
						Typename string `json:"__typename"`
						Names    struct {
							Typename string `json:"__typename"`
							Name     string `json:"name"`
						} `json:"names"`
					} `json:"metadata"`
					PrimaryImage struct {
						Typename string `json:"__typename"`
						Url      string `json:"url"`
					} `json:"primaryImage"`
					Genres []struct {
						Typename string `json:"__typename"`
						Name     string `json:"name"`
						Slug     string `json:"slug"`
					} `json:"genres"`
					ObjectRegions []struct {
						Typename string `json:"__typename"`
						Id       int    `json:"id"`
						Releases []struct {
							Typename           string `json:"__typename"`
							Id                 string `json:"id"`
							PlatformAttributes []struct {
								Typename string `json:"__typename"`
								Id       int    `json:"id"`
								Name     string `json:"name"`
								Slug     string `json:"slug"`
							} `json:"platformAttributes"`
						} `json:"releases"`
					} `json:"objectRegions"`
					Franchises []interface{} `json:"franchises"`
				} `json:"primaryObject"`
				HasAds                bool   `json:"hasAds"`
				Image                 string `json:"image"`
				DesktopInitialAdBreak int    `json:"desktopInitialAdBreak"`
				DesktopAdBreak        int    `json:"desktopAdBreak"`
				MobileInitialAdBreak  int    `json:"mobileInitialAdBreak"`
				MobileAdBreak         int    `json:"mobileAdBreak"`
				Paywall               bool   `json:"paywall"`
				WikiSlug              string `json:"wikiSlug"`
				InitialSidebarViews   struct {
					Views              Views  `json:"views"`
					InitialView        string `json:"initialView"`
					InitialSharedState struct {
						SelectedSection   string `json:"selectedSection"`
						TrackingComponent string `json:"trackingComponent"`
					} `json:"initialSharedState"`
				} `json:"initialSidebarViews"`
				RequestReferer    string `json:"requestReferer"`
				SiteEntryReferrer string `json:"siteEntryReferrer"`
				Referrer          string `json:"referrer"`
				Region            string `json:"region"`
				ParsedUrl         struct {
					Slashes  bool   `json:"slashes"`
					Protocol string `json:"protocol"`
					Hash     string `json:"hash"`
					Query    struct {
					} `json:"query"`
					Pathname string `json:"pathname"`
					Auth     string `json:"auth"`
					Host     string `json:"host"`
					Port     string `json:"port"`
					Hostname string `json:"hostname"`
					Password string `json:"password"`
					Username string `json:"username"`
					Origin   string `json:"origin"`
					Href     string `json:"href"`
				} `json:"parsedUrl"`
				UADevice string `json:"uaDevice"`
			} `json:"page"`
			PageKey string `json:"pageKey"`
			Headers struct {
				Keys []string `json:"keys"`
			} `json:"headers"`
		} `json:"pageProps"`
	} `json:"props"`
	Page  string `json:"page"`
	Query struct {
		Slug string `json:"slug"`
	} `json:"query"`
}

func (wp *WikiProps) NextPageUrl() string {
	return wp.Props.PageProps.Page.Page.NextPage.Url
}

func (wp *WikiProps) PreviousPageUrl() string {
	return wp.Props.PageProps.Page.Page.PrevPage.Url
}

func (wp *WikiProps) HTMLEntities() []HTMLEntity {
	return wp.Props.PageProps.Page.Page.HtmlEntities
}

func (he *HTMLEntity) PageUrls(slug string) ([]string, error) {
	fragment, err := html.Parse(strings.NewReader(he.Values.Html))
	if err != nil {
		return nil, err
	}

	etc := match_node.NewEtc(atom.A, "", false)

	pageUrls := make([]string, 0)

	wikiSlugPfx := fmt.Sprintf("%s/%s/", wikisLinkPfx, slug)

	for _, link := range match_node.Matches(fragment, etc, -1) {
		href := match_node.AttrVal(link, "href")
		if strings.HasPrefix(href, wikisLinkPfx) {
			rel := strings.TrimPrefix(href, wikiSlugPfx)
			if rel != "" {
				pageUrls = append(pageUrls, rel)
			}
		}
	}

	return pageUrls, nil
}

func (he *HTMLEntity) ImageUrls() ([]string, error) {
	panic("not implemented")
	//return nil, nil
}
