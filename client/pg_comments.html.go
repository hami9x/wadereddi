package client
import (
	. "fmt"
	. "strings"
	. "github.com/phaikawl/wade/utils"
	. "github.com/phaikawl/wade/core"
	. "github.com/phaikawl/wade/app/utils"
	"github.com/phaikawl/wade/dom"
)

var Tmpl_include2 = VPrep(&VNode{
	Data: "w_group",
	Type: GroupNode,	Binds: []BindFunc{
	},
	Attrs: Attributes{
		"src": "public/pg_comments.html",
		"_belong": PageComments,
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
						"class": "col-sm-1",
					},
					Children: []*VNode{
						VComponent(func() (*VNode, func(*VNode)) {
									__m := new(VoteBoxModel); __m.Init(); __node := Tmpl_component_votebox(__m)
									return __node, func(_ *VNode) {
										__m.Vote = _cvm.Post.Vote
										__m.VoteUrl = _cvm.postVoteUrl()
										__m.App = _app()
										__m.Update(__node)
									}
								}),
					},
				},
				{
					Data: "div",
					Type: ElementNode,					Attrs: Attributes{
						"class": "col-sm-11",
					},
					Children: []*VNode{
						{
							Data: "div",
							Type: ElementNode,							Children: []*VNode{
								{
									Data: "a",
									Type: ElementNode,									Binds: []BindFunc{
										func(n *VNode){ n.Attrs["href"] = ctx().getPostLink(_cvm.Post) },
									},
									Children: []*VNode{
										VMustache(func() interface{} { return _cvm.Post.Title }),
									},
								},
								{
									Data: "w_group",
									Type: GroupNode,									Binds: []BindFunc{
										func(__node *VNode) {
											__data := _cvm.Post.Labels
											__node.Children = make([]*VNode, len(__data))
											for __index, label := range __data { label := label 
												__node.Children[__index] = VPrep(&VNode{
													Data: "w_group",
													Type: GroupNode,													Children: []*VNode{
														{
															Data: "span",
															Type: ElementNode,															Attrs: Attributes{
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
							Data: "div",
							Type: ElementNode,							Children: []*VNode{
								{
									Data: "small",
									Type: ElementNode,									Attrs: Attributes{
										"class": "text-muted",
									},
									Children: []*VNode{
										VText(`submitted `),
										VMustache(func() interface{} { return _cvm.Post.Time }),
										VText(` hours ago by`),
									},
								},
								VMustache(func() interface{} { return _cvm.Post.Author }),
							},
						},
						{
							Data: "div",
							Type: ElementNode,							Binds: []BindFunc{
							},
							Attrs: Attributes{
								"class": "panel panel-default",
							},
							Children: []*VNode{
								{
									Data: "div",
									Type: ElementNode,									Attrs: Attributes{
										"class": "panel-body",
									},
									Children: []*VNode{
										VMustache(func() interface{} { return _cvm.Post.Content }),
									},
								},
							},
						},
					},
				},
			},
		},
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
						{
							Data: "div",
							Type: ElementNode,							Children: []*VNode{
								{
									Data: "small",
									Type: ElementNode,									Attrs: Attributes{
										"id": "test",
										"class": "text-muted",
									},
									Children: []*VNode{
										VMustache(func() interface{} { return len(_cvm.Comments) }),
										VText(` Comments`),
									},
								},
							},
						},
						{
							Data: "div",
							Type: ElementNode,							Children: []*VNode{
								{
									Data: "form",
									Type: ElementNode,									Children: []*VNode{
										{
											Data: "div",
											Type: ElementNode,											Attrs: Attributes{
												"class": "form-group",
											},
											Children: []*VNode{
												{
													Data: "textarea",
													Type: ElementNode,													Binds: []BindFunc{
													},
													Attrs: Attributes{
														"rows": "3",
														"cols": "80",
													},
												},
											},
										},
									},
								},
							},
						},
						{
							Data: "div",
							Type: ElementNode,							Children: []*VNode{
								{
									Data: "button",
									Type: ElementNode,									Binds: []BindFunc{
										func(n *VNode){ n.Attrs["disabled"] = _cvm.NewComment == `` },
										func(__node *VNode) {
											__node.Attrs["onclick"] = func(__event dom.Event) { _cvm.AddComment() }
										},
									},
									Attrs: Attributes{
										"class": "btn btn-success",
									},
									Children: []*VNode{
										VText(`Save`),
									},
								},
								VText(` Sort by: `),
								{
									Data: "button",
									Type: ElementNode,									Attrs: Attributes{
										"class": "btn btn-default dropdown-toggle",
										"type": "button",
										"data-toggle": "dropdown",
									},
									Children: []*VNode{
										VMustache(func() interface{} { return _cvm.RankMode }),
										{
											Data: "span",
											Type: ElementNode,											Attrs: Attributes{
												"class": "caret",
											},
										},
									},
								},
								{
									Data: "ul",
									Type: ElementNode,									Attrs: Attributes{
										"role": "menu",
										"class": "dropdown-menu",
									},
									Children: []*VNode{
										{
											Data: "w_group",
											Type: GroupNode,											Binds: []BindFunc{
												func(__node *VNode) {
													__data := _rankModes
													__node.Children = make([]*VNode, len(__data))
													for __index, mode := range __data { mode := mode 
														__node.Children[__index] = VPrep(&VNode{
															Data: "w_group",
															Type: GroupNode,															Children: []*VNode{
																{
																	Data: "li",
																	Type: ElementNode,																	Children: []*VNode{
																		{
																			Data: "a",
																			Type: ElementNode,																			Binds: []BindFunc{
																				func(__node *VNode) {
																					__node.Attrs["onclick"] = func(__event dom.Event) { _cvm.Request(mode.Code) }
																				},
																			},
																			Attrs: Attributes{
																				"href": "#",
																			},
																			Children: []*VNode{
																				VMustache(func() interface{} { return mode.Name }),
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
								},
							},
						},
					},
				},
			},
		},
		{
			Data: "w_group",
			Type: GroupNode,			Binds: []BindFunc{
				func(__node *VNode) {
					__data := _cvm.Comments
					__node.Children = make([]*VNode, len(__data))
					for __index, comment := range __data { comment := comment 
						__node.Children[__index] = VPrep(&VNode{
							Data: "w_group",
							Type: GroupNode,							Children: []*VNode{
								{
									Data: "div",
									Type: ElementNode,									Attrs: Attributes{
										"class": "row-fluid",
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
																__m.Vote = comment.Voting()
																__m.VoteUrl = _cvm.commentVoteUrl(comment)
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
													Data: "div",
													Type: ElementNode,													Children: []*VNode{
														{
															Data: "small",
															Type: ElementNode,															Attrs: Attributes{
																"class": "text-muted",
															},
															Children: []*VNode{
																VText(`submitted `),
																VMustache(func() interface{} { return comment.Time }),
																VText(` hours ago by`),
															},
														},
														VMustache(func() interface{} { return comment.Author }),
													},
												},
												{
													Data: "div",
													Type: ElementNode,													Attrs: Attributes{
														"class": "panel panel-default",
													},
													Children: []*VNode{
														{
															Data: "div",
															Type: ElementNode,															Attrs: Attributes{
																"class": "panel-body",
															},
															Children: []*VNode{
																VMustache(func() interface{} { return comment.Content }),
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