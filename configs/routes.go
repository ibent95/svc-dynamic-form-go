package configs

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

/** SetupRoutes mendaftarkan dua kelompok route:
 * 1. Web routes di path root (misal serving HTML, static files, dsb)
 * 2. API v1 routes di /public/api/v1
 */
func SetupRoutes(route *echo.Echo, opts RouteOptions) {

    /** 1) Web Routes
     * Example: serve static frontend files di folder "public/"
     */

    /** Static files if frontend assets are in "public" directory
     * Uncomment the line below to serve static files.
     * e.Static("/", "public")
     * e.GET("/", func(c echo.Context) error {
     *.....return c.File("public/index.html")
     * })
     */

	// Welcome route in root "/"
	route.GET("/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, echo.Map{
            "message": "🌐 Welcome to svc-dynamic-form-go.",
        })
    })

	// Health route in root "/healthz"
	route.GET("/healthz", func(c echo.Context) error {
        return c.JSON(http.StatusOK, echo.Map{
            "message": "✅ Service svc-dynamic-form-go is running.",
        })
    })

    // Fallback all path than root "/"
    route.HTTPErrorHandler = func(err error, c echo.Context) {
        if c.Response().Committed {
            return
        }

        if he, ok := err.(*echo.HTTPError); ok && he.Code == http.StatusNotFound {
            c.JSON(http.StatusUnauthorized, echo.Map{
                "error": "Unauthorized",
            })
        } else {
            c.JSON(http.StatusInternalServerError, echo.Map{
                "error": "Internal Server Error",
            })
        }
    }


    /** 2) API v1 Routes
	 * - Group : /public/api/v1
     * - Uncomment the lines below to enable API routes. for example:
     *      routeApi.GET("/users", opts.UserController.Index)
     * - Add new group if necessary, for example:
     *      admin := route.Group("/admin")
     *      admin.Use(middleware.JWTWithConfig(...))
     *      admin.GET("/users", controllers.GetUsers)
     */
    apiRoute := route.Group("/public/api/v1")

    apiRoute.GET("/organization", opts.OrganizationController.Index)
    apiRoute.POST("/organization", opts.OrganizationController.Store)
    apiRoute.GET("/publication", opts.PublicationController.Index)
    apiRoute.GET("/user", opts.UserController.Index)

}
