package client
import (
	. "fmt"
	. "strings"
	. "github.com/phaikawl/wade/utils"
	. "github.com/phaikawl/wade/core"
	. "github.com/phaikawl/wade/app/utils"
	"github.com/phaikawl/wade/dom"
)

var Tmpl_main = VPrep(&VNode{
	Data: "w_group",
	Type: GroupNode,	Children: []*VNode{
		{
			Data: "div",
			Type: ElementNode,			Attrs: Attributes{
				"class": "wrapper",
			},
			Children: []*VNode{
				{
					Data: "div",
					Type: ElementNode,					Attrs: Attributes{
						"class": "box",
					},
					Children: []*VNode{
						{
							Data: "div",
							Type: ElementNode,							Attrs: Attributes{
								"class": "row",
							},
							Children: []*VNode{
								{
									Data: "div",
									Type: ElementNode,									Attrs: Attributes{
										"class": "column col-sm-2",
										"id": "sidebar",
									},
									Children: []*VNode{
										{
											Data: "ul",
											Type: ElementNode,											Attrs: Attributes{
												"class": "nav",
											},
											Children: []*VNode{
												{
													Data: "li",
													Type: ElementNode,													Children: []*VNode{
														{
															Data: "a",
															Type: ElementNode,															Binds: []BindFunc{
																func(n *VNode){ n.Attrs["href"] = Url(PagePosts, `top`) },
															},
															Children: []*VNode{
																VText(`Posts`),
															},
														},
													},
												},
											},
										},
										{
											Data: "ul",
											Type: ElementNode,											Attrs: Attributes{
												"class": "nav hidden-xs",
												"id": "sidebar-footer",
											},
											Children: []*VNode{
												{
													Data: "li",
													Type: ElementNode,													Children: []*VNode{
														{
															Data: "a",
															Type: ElementNode,															Attrs: Attributes{
																"href": "#",
															},
															Children: []*VNode{
																{
																	Data: "h3",
																	Type: ElementNode,																	Children: []*VNode{
																		VText(`WadeReddi`),
																	},
																},
																VText(`From Hai with `),
																{
																	Data: "i",
																	Type: ElementNode,																	Attrs: Attributes{
																		"class": "glyphicon glyphicon-heart-empty",
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
									Type: ElementNode,									Attrs: Attributes{
										"class": "column col-sm-10",
										"id": "main",
									},
									Children: []*VNode{
										{
											Data: "div",
											Type: ElementNode,											Attrs: Attributes{
												"class": "padding",
											},
											Children: []*VNode{
												{
													Data: "div",
													Type: ElementNode,													Attrs: Attributes{
														"class": "full col-sm-9",
													},
													Children: []*VNode{
														Tmpl_include1,
														Tmpl_include2,
														{
															Data: "div",
															Type: ElementNode,															Binds: []BindFunc{
															},
															Children: []*VNode{
																VText(` We are sorry, no such thing is here. `),
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

func init() {_ = Url; _ = Join; _ = ToString; _ = Sprintf; _ = dom.DebugInfo}