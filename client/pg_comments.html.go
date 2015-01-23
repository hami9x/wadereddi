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
	Type: GroupNode,
	Binds: []BindFunc{
	},
	Attrs: Attributes{
		"src": "public/pg_comments.html",
		"_belong": PageComments,
	},
	Children: []*VNode{
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
						"class": "col-sm-1",
					},
					Children: []*VNode{
						VText(` `),
						VComponent(func() (*VNode, func(*VNode)) {
									__m := new(VoteBoxModel); __m.Init(); __node := Tmpl_component_votebox(__m)
									return __node, func(_ *VNode) {
										__m.Vote = _cvm.Post.Vote
										__m.VoteUrl = _cvm.postVoteUrl()
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
							Data: "div",
							Type: ElementNode,
							Children: []*VNode{
								VText(` `),
								{
									Data: "a",
									Type: ElementNode,
									Binds: []BindFunc{
										func(n *VNode){ n.Attrs["href"] = ctx().getPostLink(_cvm.Post) },
									},
									Children: []*VNode{
										VMustache(func() interface{} { return _cvm.Post.Title }),
									},
								},
								VText(` `),
								{
									Data: "w_group",
									Type: GroupNode,
									Binds: []BindFunc{
										func(__node *VNode) {
											__data := _cvm.Post.Labels
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
							Data: "div",
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
										VText(`submitted `),
										VMustache(func() interface{} { return _cvm.Post.Time }),
										VText(` hours ago by`),
									},
								},
								VText(` `),
								VMustache(func() interface{} { return _cvm.Post.Author }),
								VText(` `),
							},
						},
						VText(` `),
						{
							Data: "div",
							Type: ElementNode,
							Binds: []BindFunc{
							},
							Attrs: Attributes{
								"class": "panel panel-default",
							},
							Children: []*VNode{
								VText(` `),
								{
									Data: "div",
									Type: ElementNode,
									Attrs: Attributes{
										"class": "panel-body",
									},
									Children: []*VNode{
										VMustache(func() interface{} { return _cvm.Post.Content }),
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
						{
							Data: "div",
							Type: ElementNode,
							Children: []*VNode{
								VText(` `),
								{
									Data: "small",
									Type: ElementNode,
									Attrs: Attributes{
										"class": "text-muted",
										"id": "test",
									},
									Children: []*VNode{
										VMustache(func() interface{} { return len(_cvm.Comments) }),
										VText(` Comments`),
									},
								},
								VText(` `),
							},
						},
						VText(` `),
						{
							Data: "div",
							Type: ElementNode,
							Children: []*VNode{
								VText(` `),
								{
									Data: "form",
									Type: ElementNode,
									Children: []*VNode{
										VText(` `),
										{
											Data: "div",
											Type: ElementNode,
											Attrs: Attributes{
												"class": "form-group",
											},
											Children: []*VNode{
												VText(` `),
												{
													Data: "textarea",
													Type: ElementNode,
													Binds: []BindFunc{
													},
													Attrs: Attributes{
														"rows": "3",
														"cols": "80",
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
						{
							Data: "div",
							Type: ElementNode,
							Children: []*VNode{
								VText(` `),
								{
									Data: "button",
									Type: ElementNode,
									Binds: []BindFunc{
										func(__node *VNode) {
											__node.Attrs["onclick"] = func(__event dom.Event) { __event.PreventDefault(); _cvm.AddComment() }
										},
										func(n *VNode){ n.Attrs["disabled"] = _cvm.NewComment == `` },
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
									Type: ElementNode,
									Attrs: Attributes{
										"class": "btn btn-default dropdown-toggle",
										"type": "button",
										"data-toggle": "dropdown",
									},
									Children: []*VNode{
										VText(` `),
										VMustache(func() interface{} { return _cvm.RankMode }),
										VText(` `),
										{
											Data: "span",
											Type: ElementNode,
											Attrs: Attributes{
												"class": "caret",
											},
										},
										VText(` `),
									},
								},
								VText(` `),
								{
									Data: "ul",
									Type: ElementNode,
									Attrs: Attributes{
										"class": "dropdown-menu",
										"role": "menu",
									},
									Children: []*VNode{
										VText(` `),
										{
											Data: "w_group",
											Type: GroupNode,
											Binds: []BindFunc{
												func(__node *VNode) {
													__data := _rankModes
													__node.Children = make([]*VNode, len(__data))
													for __index, __value := range __data { mode := __value 
														__node.Children[__index] = VPrep(&VNode{
															Data: "w_group",
															Type: GroupNode,
															Children: []*VNode{
																VText(` `),
																{
																	Data: "li",
																	Type: ElementNode,
																	Children: []*VNode{
																		VText(` `),
																		{
																			Data: "a",
																			Type: ElementNode,
																			Binds: []BindFunc{
																				func(__node *VNode) {
																					__node.Attrs["onclick"] = func(__event dom.Event) { __event.PreventDefault(); go _cvm.Request(mode.Code) }
																				},
																			},
																			Attrs: Attributes{
																				"href": "#",
																			},
																			Children: []*VNode{
																				VText(` `),
																				VMustache(func() interface{} { return mode.Name }),
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
		},
		VText(` `),
		{
			Data: "w_group",
			Type: GroupNode,
			Binds: []BindFunc{
				func(__node *VNode) {
					__data := _cvm.Comments
					__node.Children = make([]*VNode, len(__data))
					for __index, __value := range __data { comment := __value 
						__node.Children[__index] = VPrep(&VNode{
							Data: "w_group",
							Type: GroupNode,
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
												"class": "col-sm-1",
											},
											Children: []*VNode{
												VText(` `),
												VComponent(func() (*VNode, func(*VNode)) {
															__m := new(VoteBoxModel); __m.Init(); __node := Tmpl_component_votebox(__m)
															return __node, func(_ *VNode) {
																__m.Vote = comment.Voting()
																__m.VoteUrl = _cvm.commentVoteUrl(comment)
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
													Data: "div",
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
																VText(`submitted `),
																VMustache(func() interface{} { return comment.Time }),
																VText(` hours ago by`),
															},
														},
														VText(` `),
														VMustache(func() interface{} { return comment.Author }),
														VText(` `),
													},
												},
												VText(` `),
												{
													Data: "div",
													Type: ElementNode,
													Attrs: Attributes{
														"class": "panel panel-default",
													},
													Children: []*VNode{
														VText(` `),
														{
															Data: "div",
															Type: ElementNode,
															Attrs: Attributes{
																"class": "panel-body",
															},
															Children: []*VNode{
																VText(` `),
																VMustache(func() interface{} { return comment.Content }),
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