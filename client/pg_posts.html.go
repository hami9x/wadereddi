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
	Type: GroupNode,
	Binds: []BindFunc{
	},
	Attrs: Attributes{
		"src": "public/pg_posts.html",
		"_belong": PagePosts,
	},
	Children: []*VNode{
		VText(` `),
		{
			Data: "div",
			Type: ElementNode,
			Attrs: Attributes{
				"class": "row-fluid",
			},
			Children: []*VNode{
				VText(` `),
				{
					Data: "div",
					Type: ElementNode,
					Attrs: Attributes{
						"class": "col-sm-12",
					},
					Children: []*VNode{
						VText(` `),
						VText(` `),
					},
				},
				VText(` `),
			},
		},
		VText(` `),
		{
			Data: "w_group",
			Type: GroupNode,
			Binds: []BindFunc{
				func(__node *VNode) {
					__data := _pvm.Posts
					__node.Children = make([]*VNode, len(__data))
					for __index, __value := range __data { post := __value 
						__node.Children[__index] = VPrep(&VNode{
							Data: "w_group",
							Type: GroupNode,
							Children: []*VNode{
								VText(` `),
								{
									Data: "div",
									Type: ElementNode,
									Attrs: Attributes{
										"class": "row-fluid post-wrapper",
									},
									Children: []*VNode{
										VText(` `),
										{
											Data: "div",
											Type: ElementNode,
											Attrs: Attributes{
												"class": "col-sm-1",
											},
											Children: []*VNode{
												VText(` `),
												VComponent(func() (*VNode, func(*VNode)) {
															__m := new(VoteBoxModel); __m.Init(); __node := Tmpl_component_votebox(__m)
															return __node, func(_ *VNode) {
																__m.Vote = post.Vote
																__m.VoteUrl = _pvm.voteUrl(post)
																__m.App = _app()
																__m.Update(__node)
															}
														}),
												VText(` `),
											},
										},
										VText(` `),
										{
											Data: "div",
											Type: ElementNode,
											Attrs: Attributes{
												"class": "col-sm-11",
											},
											Children: []*VNode{
												VText(` `),
												{
													Data: "h3",
													Type: ElementNode,
													Children: []*VNode{
														{
															Data: "a",
															Type: ElementNode,
															Binds: []BindFunc{
																func(n *VNode){ n.Attrs["href"] = ctx().getPostLink(post) },
															},
															Children: []*VNode{
																VMustache(func() interface{} { return post.Title }),
															},
														},
													},
												},
												VText(` `),
												{
													Data: "h4",
													Type: ElementNode,
													Children: []*VNode{
														VText(` by `),
														VMustache(func() interface{} { return post.Author }),
														VText(` `),
														{
															Data: "w_group",
															Type: GroupNode,
															Binds: []BindFunc{
																func(__node *VNode) {
																	__data := post.Labels
																	__node.Children = make([]*VNode, len(__data))
																	for __index, __value := range __data { label := __value 
																		__node.Children[__index] = VPrep(&VNode{
																			Data: "w_group",
																			Type: GroupNode,
																			Children: []*VNode{
																				VText(` `),
																				{
																					Data: "span",
																					Type: ElementNode,
																					Attrs: Attributes{
																						"class": "label label-default",
																					},
																					Children: []*VNode{
																						VMustache(func() interface{} { return label }),
																					},
																				},
																				VText(` `),
																			},
																		})
																	}
																},
															},
															Children: []*VNode{
															},
														},
														VText(` `),
													},
												},
												VText(` `),
												{
													Data: "h4",
													Type: ElementNode,
													Children: []*VNode{
														VText(` `),
														{
															Data: "small",
															Type: ElementNode,
															Attrs: Attributes{
																"class": "text-muted",
															},
															Children: []*VNode{
																VMustache(func() interface{} { return post.Time }),
																VText(` hours ago
										• `),
																{
																	Data: "a",
																	Type: ElementNode,
																	Binds: []BindFunc{
																		func(n *VNode){ n.Attrs["href"] = Url(PageComments, post.Id) },
																	},
																	Children: []*VNode{
																		VText(` `),
																		VMustache(func() interface{} { return len(post.Comments) }),
																		VText(` Comments `),
																	},
																},
															},
														},
														VText(` `),
													},
												},
												VText(` `),
											},
										},
										VText(` `),
									},
								},
								VText(` `),
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