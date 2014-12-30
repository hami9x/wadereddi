package client
import (
	 . "github.com/phaikawl/wade/app/helpers"
	 wc "github.com/phaikawl/wade/core"
)

var include3 = wc.VPrep(wc.VNode{
	Data: "w_group",
	Children: []wc.VNode{
		{
			Data: "div",
			Attrs: wc.Attributes{
				"class": "wrapper",			
			},
			Children: []wc.VNode{
				{
					Data: "div",
					Attrs: wc.Attributes{
						"class": "box",					
					},
					Children: []wc.VNode{
						{
							Data: "div",
							Attrs: wc.Attributes{
								"class": "row",							
							},
							Children: []wc.VNode{
								{
									Data: "div",
									Attrs: wc.Attributes{
										"id": "sidebar",
										"class": "column col-sm-2",									
									},
									Children: []wc.VNode{
										{
											Data: "ul",
											Attrs: wc.Attributes{
												"class": "nav",											
											},
											Children: []wc.VNode{
												{
													Data: "li",
													Children: []wc.VNode{
														{
															Data: "a",
															Attrs: wc.Attributes{
																"@href": "url(PagePosts, `top`)",															
															},
															Children: []wc.VNode{
																wc.VText(`Posts`),
															},
														},
													},
												},
											},
										},
										{
											Data: "ul",
											Attrs: wc.Attributes{
												"class": "nav hidden-xs",
												"id": "sidebar-footer",											
											},
											Children: []wc.VNode{
												{
													Data: "li",
													Children: []wc.VNode{
														{
															Data: "a",
															Attrs: wc.Attributes{
																"href": "#",															
															},
															Children: []wc.VNode{
																{
																	Data: "h3",
																	Children: []wc.VNode{
																		wc.VText(`WadeReddi`),
																	},
																},
																wc.VText(`From Hai with `),
																VElem("i", "glyphicon glyphicon-heart-empty"),
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
										"class": "column col-sm-10",
										"id": "main",									
									},
									Children: []wc.VNode{
										{
											Data: "div",
											Attrs: wc.Attributes{
												"class": "padding",											
											},
											Children: []wc.VNode{
												{
													Data: "div",
													Attrs: wc.Attributes{
														"class": "full col-sm-9",													
													},
													Children: []wc.VNode{
														include1,
														include2,
														{
															Data: "div",
															Attrs: wc.Attributes{
																"@_belong": "PageNotFound",															
															},
															Children: []wc.VNode{
																wc.VText(` We are sorry, no such thing is here. `),
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
	},
})
