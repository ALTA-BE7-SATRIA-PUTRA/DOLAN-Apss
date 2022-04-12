package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string      `gorm:"not null" json:"name" form:"name"`
	City      string      `gorm:"not null" json:"city" form:"city"`
	UrlImage  string      `gorm:"not null;default:data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAOEAAADhCAMAAAAJbSJIAAAAP1BMVEXk5ueutLfn6eqrsbTg4uPp6+zj5eayuLu2u77Jzc+rsbXQ09XGyszU19mvtbjm6Ojb3d66v8LHysy6vsG/xMWYMMSFAAAFWUlEQVR4nO2d6XKrMAyFQcYYzN7w/s967XBJkzQL2HIsE33Tmf7l9MiSVzXLGIZhGIZhGIZhGIZhGIZhGIZhGIZhGIZhGIZhGIY5ICBlOXZdY+nGUkqI/UWYgCy7n0oLYX4s9lfVdoWEIfanoSCLujKi8juMSlWPkLyVAE0l/qj7VamaLGmNkNX5U3mryLZMVyPUz+270qjbRH2UndDv9S0+NjL21+4HsmqDfxeNqkjNRmPgdn1njYnZCO3GAP1FT0m5uCdCLy72ZezP3ky5X96icUzExsLBwKQkFruHYGIS3R1MRWLvI9BILGILeANUfgLzvI8t4TVy9orRs4kT5dIPnUeWWdE14aFYejtoEWNsHU+R3oNwoadqIjQoFhoTW6oSkQSSLRnQoinMaeZTv8nMLZpispETnsA8rwiaiGkhyYoB/rOZG070TMQVmGty6RSrFq6Qq4lS4Qo0EomFqc/C/onCLramW2rkIDXMpMIUP0iphSnOsulOIamSuHcPf5NCSithaPEF0ir6IYahgZKHAYKU1ioRd9Z9UUioIoZINKRSDdToM5ozdGr+MAcRSGgvA1CX978oMh76H1Y8hs6+KYQph3keW9iFUApFbGEXBs8zQ1YYn+NH6RdkmuNXi+NX/OPP2uowCunMvNE3vBcIrZ6+YAUcZheD0uHM8XeiwiyBSV0aPvyO8Bfs6gc5maFkoX0dg66QUL234FdETagaWo5/yo12a+8ikNpNBfRsSulYZgH7xhDBa1/It76I5RkL7s09Rc9CXBNpzWdWMEcipSP8a9BOEekl0oUvuMmONrGh22oBJ9nQTDMLKHEqfqjGqAXhZRfF2cw1g/9QpDsIF1xfOa9QLRRX+GUbncArWS+JuktAYAYer9WTEGiyjeu19hQecq8oB406T6m/icM6Q0zUy8Qtcu/shvT75odAsSdSRZ9ShK4YG7f6KGraM7VnALR/27Q90jcn2gkrs82ift75KMScYoD+MmSNetGRTqs6Xf9WAIofpR+oFKJvD9BV8AzIopsrcYXWam6Kg8j7D8ihGLumNjTdWMCxunvek9bkZSMAIG+Ag4So0QWljc52nqdKqb7P+14pVU1ze47WzGqN/ZVODNajsWunKl9yy4NUapvS5v1pbsywTEwmwDA2sxKPasTDoi/6U9ulItPkzK6tTDnYou1O5tQU1FPsIMf6RVPdt5i/y9xlZGfhAGO7MS5fm1kZK2OL+YuRN/uruxJZ0hIpi7bHvaggxNSRSTwA3YRm37XIvibRJXqAxmVjbZtGMY+xg3WANoR9v+gqqkYYAuuzCBVN40f0nTVG8hGaTdtMKOjp83s5MAbLLw8R7fBRjVAi39PbIDH/5LmUbLZ2zMdET5/aloPsFEGf5UOd6Xd3zEfkI53pYY5k4BmRhz5EhULFc3DRGPYQB8bI+qzEkJG6+9gzCEKFE4jYZNYLEWiGI6eYOeaGMJc2gI7AMBLhRCREF/DvFhETiC9RTsQE2kDFPMLCuBiLDuY9zRDvJzHAExhxrv0StM4gYTomICBmnDlqsP46/ogGw0WSWWYFI9uQHYQLGK/4SAtEeH6C+24yAN4NUAgsed/g+cQGAvWaQ8Qvn0JDaMX0FC+FsT9+Cz5dXqjOR+85uIU+JqL3ugiFc/fBBBLpgnOjF/q1cMWxJobqhhgA1wl4Mha65hrai4pbnJYYCQWpMdHl30HLhCw0k1MHhWR3Zx7icuSWSrn/z/6BmNQwdFoIB2rYGQqHvlJhmq6GY39FTCvRuHRZTKneW6q9CpNZOa3sXkENtUgKTee/DDAMwzAMwzAMwzAMwzBfyj+xS1VU5WkGtAAAAABJRU5ErkJggg==" json:"url_image" form:"url_image"`
	Email     string      `gorm:"unique;not null" json:"email" form:"email"`
	Password  string      `gorm:"not null" json:"password" form:"password"`
	Event     []Event     `gorm:"foreignKey:UserId;references:ID"`
	Attendees []Attendees `gorm:"foreignKey:UserId;references:ID"`
	Comment   []Comment   `gorm:"foreignKey:UserId;references:ID"`
}
