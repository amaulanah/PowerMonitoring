type DynamicRoutes = {
	"/pm/[meterId]": { meterId: string }
};

type Layouts = {
	"/": { meterId?: string };
	"/login": undefined;
	"/pm": { meterId?: string };
	"/pm/[meterId]": { meterId: string }
};

export type RouteId = "/" | "/login" | "/pm" | "/pm/[meterId]";

export type RouteParams<T extends RouteId> = T extends keyof DynamicRoutes ? DynamicRoutes[T] : Record<string, never>;

export type LayoutParams<T extends RouteId> = Layouts[T] | Record<string, never>;

export type Pathname = "/" | "/login" | "/pm" | `/pm/${string}` & {};

export type ResolvedPathname = `${"" | `/${string}`}${Pathname}`;

export type Asset = "/favicon.png";