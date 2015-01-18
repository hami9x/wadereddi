package client
import (
	. "fmt"
	. "strings"
	. "github.com/phaikawl/wade/utils"
	. "github.com/phaikawl/wade/core"
	. "github.com/phaikawl/wade/app/utils"
	"github.com/phaikawl/wade/dom"
)

var Tmpl_component_votebox = func(__m *VoteBoxModel) *VNode {
	return VPrep(&VNode{
			Data: "votebox",
			Type: ElementNode,			Children: []*VNode{
				{
					Data: "div",
					Type: ElementNode,					Attrs: Attributes{
						"class": "votebox",
					},
					Children: []*VNode{
						{
							Data: "div",
							Type: ElementNode,							Attrs: Attributes{
								"class": "row-fluid",
							},
						},
						{
							Data: "div",
							Type: ElementNode,							Attrs: Attributes{
								"class": "row-fluid",
							},
							Children: []*VNode{
								{
									Data: "a",
									Type: ElementNode,									Binds: []BindFunc{
										func(__node *VNode) {
											__node.Attrs["onclick"] = func(__event dom.Event) { __m.DoVote(1) }
										},
									},
									Attrs: Attributes{
										"class": "upvote-btn",
										"href": "#",
									},
									Children: []*VNode{
										{
											Data: "i",
											Type: ElementNode,											Attrs: Attributes{
												"class": "glyphicon glyphicon-arrow-up",
											},
										},
									},
								},
							},
						},
						{
							Data: "div",
							Type: ElementNode,							Attrs: Attributes{
								"class": "row-fluid score",
							},
							Children: []*VNode{
								VMustache(func() interface{} { return __m.Vote.Score }),
							},
						},
						{
							Data: "div",
							Type: ElementNode,							Attrs: Attributes{
								"class": "row-fluid",
							},
							Children: []*VNode{
								{
									Data: "a",
									Type: ElementNode,									Binds: []BindFunc{
										func(__node *VNode) {
											__node.Attrs["onclick"] = func(__event dom.Event) { __m.DoVote(-1) }
										},
									},
									Attrs: Attributes{
										"class": "downvote-btn",
										"href": "#",
									},
									Children: []*VNode{
										{
											Data: "i",
											Type: ElementNode,											Attrs: Attributes{
												"class": "glyphicon glyphicon-arrow-down",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		})
}

func init() {_ = Url; _ = Join; _ = ToString; _ = Sprintf; _ = dom.DebugInfo}