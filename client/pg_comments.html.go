package client
import (
	 . "github.com/phaikawl/wade/app/helpers"
	 wc "github.com/phaikawl/wade/core"
)

var include2 = wc.VPrep(wc.VNode{
	Data: "w_group",
	Attrs: wc.Attributes{
		"src": "public/pg_comments.html",	
	},
	Children: []wc.VNode{
		{
			Data: "div",
			Attrs: wc.Attributes{
				"class": "row-fluid",			
			},
			Children: []wc.VNode{
				{
					Data: "div",
					Attrs: wc.Attributes{
						"class": "col-sm-1",					
					},
					Children: []wc.VNode{
						{
							Data: "c:votebox",
							Attrs: wc.Attributes{
								"*vote": "_cvm.Post.Vote",
								"*vote_url": "_cvm.postVoteUrl()",							
							},
						},
					},
				},
				{
					Data: "div",
					Attrs: wc.Attributes{
						"class": "col-sm-11",					
					},
					Children: []wc.VNode{
						{
							Data: "div",
							Children: []wc.VNode{
								{
									Data: "a",
									Attrs: wc.Attributes{
										"@href": "ctx().getLink(_cvm.Post)",									
									},
									Children: []wc.VNode{
										wc.VMustache(func() interface{} { return  _cvm.Post.Title  }),
									},
								},
								{
									Data: "w_group",
									Attrs: wc.Attributes{
										"#range(_,label)": "_cvm.Post.Labels",									
									},
									Children: []wc.VNode{
										{
											Data: "span",
											Attrs: wc.Attributes{
												"class": "label label-default",											
											},
											Children: []wc.VNode{
												wc.VMustache(func() interface{} { return  label  }),
											},
										},
									},
								},
							},
						},
						{
							Data: "div",
							Children: []wc.VNode{
								{
									Data: "small",
									Attrs: wc.Attributes{
										"class": "text-muted",									
									},
									Children: []wc.VNode{
										wc.VText(`submitted `),
										wc.VMustache(func() interface{} { return  _cvm.Post.Time  }),
										wc.VText(` hours ago by`),
									},
								},
								wc.VMustache(func() interface{} { return  _cvm.Post.Author  }),
							},
						},
						{
							Data: "div",
							Attrs: wc.Attributes{
								"class": "panel panel-default",
								"#ifn": "_cvm.Post.Content == ``",							
							},
							Children: []wc.VNode{
								{
									Data: "div",
									Attrs: wc.Attributes{
										"class": "panel-body",									
									},
									Children: []wc.VNode{
										wc.VMustache(func() interface{} { return  _cvm.Post.Content  }),
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
			Attrs: wc.Attributes{
				"class": "row-fluid",			
			},
			Children: []wc.VNode{
				{
					Data: "div",
					Attrs: wc.Attributes{
						"class": "col-sm-12",					
					},
					Children: []wc.VNode{
						{
							Data: "div",
							Children: []wc.VNode{
								{
									Data: "small",
									Attrs: wc.Attributes{
										"class": "text-muted",
										"id": "test",									
									},
									Children: []wc.VNode{
										wc.VMustache(func() interface{} { return  len(_cvm.Comments)  }),
										wc.VText(` Comments`),
									},
								},
							},
						},
						{
							Data: "div",
							Children: []wc.VNode{
								{
									Data: "form",
									Children: []wc.VNode{
										{
											Data: "div",
											Attrs: wc.Attributes{
												"class": "form-group",											
											},
											Children: []wc.VNode{
												{
													Data: "textarea",
													Attrs: wc.Attributes{
														"rows": "3",
														"cols": "80",
														"#value(change)": "_cvm.NewComment",													
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
							Children: []wc.VNode{
								{
									Data: "button",
									Attrs: wc.Attributes{
										"#on(click)": "_cvm.AddComment",
										"class": "btn btn-success",
										"@disabled": "_cvm.NewComment == ``",									
									},
									Children: []wc.VNode{
										wc.VText(`Save`),
									},
								},
								wc.VText(` Sort by: `),
								{
									Data: "button",
									Attrs: wc.Attributes{
										"type": "button",
										"data-toggle": "dropdown",
										"class": "btn btn-default dropdown-toggle",									
									},
									Children: []wc.VNode{
										wc.VMustache(func() interface{} { return  _cvm.RankMode  }),
										VElem("span", "caret"),
									},
								},
								{
									Data: "ul",
									Attrs: wc.Attributes{
										"class": "dropdown-menu",
										"role": "menu",									
									},
									Children: []wc.VNode{
										{
											Data: "li",
											Attrs: wc.Attributes{
												"#range(_,mode)": "_rankModes",											
											},
											Children: []wc.VNode{
												{
													Data: "a",
													Attrs: wc.Attributes{
														"href": "#",
														"#on(click)": "func(){ _cvm.Request(mode.Code) }",													
													},
													Children: []wc.VNode{
														wc.VMustache(func() interface{} { return  mode.Name  }),
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
		},
		{
			Data: "div",
			Attrs: wc.Attributes{
				"#range(_,comment)": "_cvm.Comments",
				"class": "row-fluid",			
			},
			Children: []wc.VNode{
				{
					Data: "div",
					Attrs: wc.Attributes{
						"class": "col-sm-1",					
					},
					Children: []wc.VNode{
						{
							Data: "c:votebox",
							Attrs: wc.Attributes{
								"*vote": "comment.Voting()",
								"*vote_url": "_cvm.commentVoteUrl(comment)",							
							},
						},
					},
				},
				{
					Data: "div",
					Attrs: wc.Attributes{
						"class": "col-sm-11",					
					},
					Children: []wc.VNode{
						{
							Data: "div",
							Children: []wc.VNode{
								{
									Data: "small",
									Attrs: wc.Attributes{
										"class": "text-muted",									
									},
									Children: []wc.VNode{
										wc.VText(`submitted `),
										wc.VMustache(func() interface{} { return  comment.Time  }),
										wc.VText(` hours ago by`),
									},
								},
								wc.VMustache(func() interface{} { return  comment.Author  }),
							},
						},
						{
							Data: "div",
							Attrs: wc.Attributes{
								"class": "panel panel-default",							
							},
							Children: []wc.VNode{
								{
									Data: "div",
									Attrs: wc.Attributes{
										"class": "panel-body",									
									},
									Children: []wc.VNode{
										wc.VMustache(func() interface{} { return  comment.Content  }),
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
