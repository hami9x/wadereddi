package client
import (
	 . "github.com/phaikawl/wade/app/helpers"
	 wc "github.com/phaikawl/wade/core"
)

var include1 = wc.VPrep(wc.VNode{
	Data: "w_group",
	Attrs: wc.Attributes{
		"src": "public/pg_posts.html",	
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
						"class": "col-sm-12",					
					},
					Children: []wc.VNode{
						{
							Data: "c:switch_menu",
							Attrs: wc.Attributes{
								"*current": "_pvm.RankMode",							
							},
							Children: []wc.VNode{
								{
									Data: "ul",
									Attrs: wc.Attributes{
										"class": "nav nav-pills",									
									},
									Children: []wc.VNode{
										{
											Data: "w_group",
											Attrs: wc.Attributes{
												"#range(_,mode)": "_rankModes",											
											},
											Children: []wc.VNode{
												{
													Data: "li",
													Attrs: wc.Attributes{
														"@case": "mode.Code",													
													},
													Children: []wc.VNode{
														{
															Data: "a",
															Attrs: wc.Attributes{
																"@href": "url(PagePosts, mode.Code)",															
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
			},
		},
		{
			Data: "div",
			Attrs: wc.Attributes{
				"#range(_,post)": "_pvm.Posts",
				"class": "row-fluid post-wrapper",			
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
								"*vote": "post.Vote",
								"*vote_url": "_pvm.voteUrl(post)",							
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
							Data: "h3",
							Children: []wc.VNode{
								{
									Data: "a",
									Attrs: wc.Attributes{
										"@href": "ctx().getLink(post)",									
									},
									Children: []wc.VNode{
										wc.VMustache(func() interface{} { return  post.Title  }),
									},
								},
							},
						},
						{
							Data: "h4",
							Children: []wc.VNode{
								wc.VText(` by `),
								wc.VMustache(func() interface{} { return  post.Author  }),
								{
									Data: "span",
									Attrs: wc.Attributes{
										"class": "label label-default",
										"#range(_,label)": "post.Labels",									
									},
									Children: []wc.VNode{
										wc.VMustache(func() interface{} { return  label  }),
									},
								},
							},
						},
						{
							Data: "h4",
							Children: []wc.VNode{
								{
									Data: "small",
									Attrs: wc.Attributes{
										"class": "text-muted",									
									},
									Children: []wc.VNode{
										wc.VMustache(func() interface{} { return  post.Time  }),
										wc.VText(` hours ago
									• `),
										{
											Data: "a",
											Attrs: wc.Attributes{
												"@href": "url(PageComments, post.Id)",											
											},
											Children: []wc.VNode{
												wc.VMustache(func() interface{} { return  len(post.Comments)  }),
												wc.VText(` Comments `),
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
