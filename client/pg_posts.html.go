package client
import (
	. "fmt"
	. "strings"
	. "github.com/phaikawl/wade/utils"
	. "github.com/phaikawl/wade/core"
	. "github.com/phaikawl/wade/app/utils"
	"github.com/phaikawl/wade/dom"
)

var Tmpl_include1 = VPrep(&VNode{
	Data: "w_group",
	Type: GroupNode,	Binds: []BindFunc{
	},
	Attrs: Attributes{
		"src": "public/pg_posts.html",
		"_belong": PagePosts,
	},
	Children: []*VNode{
		{
			Data: "div",
			Type: ElementNode,			Attrs: Attributes{
				"class": "row-fluid",
			},
			Children: []*VNode{
				{
					Data: "div",
					Type: ElementNode,					Attrs: Attributes{
						"class": "col-sm-12",
					},
					Children: []*VNode{
					},
				},
			},
		},
		{
			Data: "w_group",
			Type: GroupNode,			Binds: []BindFunc{
				func(__node *VNode) {
					__data := _pvm.Posts
					__node.Children = make([]*VNode, len(__data))
					for __index, post := range __data { post := post 
						__node.Children[__index] = VPrep(&VNode{
							Data: "w_group",
							Type: GroupNode,							Children: []*VNode{
								{
									Data: "div",
									Type: ElementNode,									Attrs: Attributes{
										"class": "row-fluid post-wrapper",
									},
									Children: []*VNode{
										{
											Data: "div",
											Type: ElementNode,											Attrs: Attributes{
												"class": "col-sm-1",
											},
											Children: []*VNode{
												VComponent(func() (*VNode, func(*VNode)) {
															__m := new(VoteBoxModel); __m.Init(); __node := Tmpl_component_votebox(__m)
															return __node, func(_ *VNode) {
																__m.Vote = post.Vote
																__m.VoteUrl = _pvm.voteUrl(post)
																__m.App = _app()
																__m.Update(__node)
															}
														}),
											},
										},
										{
											Data: "div",
											Type: ElementNode,											Attrs: Attributes{
												"class": "col-sm-11",
											},
											Children: []*VNode{
												{
													Data: "h3",
													Type: ElementNode,													Children: []*VNode{
														{
															Data: "a",
															Type: ElementNode,															Binds: []BindFunc{
																func(n *VNode){ n.Attrs["href"] = ctx().getPostLink(post) },
															},
															Children: []*VNode{
																VMustache(func() interface{} { return post.Title }),
															},
														},
													},
												},
												{
													Data: "h4",
													Type: ElementNode,													Children: []*VNode{
														VText(` by `),
														VMustache(func() interface{} { return post.Author }),
														{
															Data: "w_group",
															Type: GroupNode,															Binds: []BindFunc{
																func(__node *VNode) {
																	__data := post.Labels
																	__node.Children = make([]*VNode, len(__data))
																	for __index, label := range __data { label := label 
																		__node.Children[__index] = VPrep(&VNode{
																			Data: "w_group",
																			Type: GroupNode,																			Children: []*VNode{
																				{
																					Data: "span",
																					Type: ElementNode,																					Attrs: Attributes{
																						"class": "label label-default",
																					},
																					Children: []*VNode{
																						VMustache(func() interface{} { return label }),
																					},
																				},
																			},
																		})
																	}
																},
															},
															Children: []*VNode{
															},
														},
													},
												},
												{
													Data: "h4",
													Type: ElementNode,													Children: []*VNode{
														{
															Data: "small",
															Type: ElementNode,															Attrs: Attributes{
																"class": "text-muted",
															},
															Children: []*VNode{
																VMustache(func() interface{} { return post.Time }),
																VText(` hours ago
										• `),
																{
																	Data: "a",
																	Type: ElementNode,																	Binds: []BindFunc{
																		func(n *VNode){ n.Attrs["href"] = Url(PageComments, post.Id) },
																	},
																	Children: []*VNode{
																		VMustache(func() interface{} { return len(post.Comments) }),
																		VText(` Comments `),
																	},
																},
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
				},
			},
			Children: []*VNode{
			},
		},
	},
})

func init() {_ = Url; _ = Join; _ = ToString; _ = Sprintf; _ = dom.DebugInfo}