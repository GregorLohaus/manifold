import { env } from "$env/dynamic/public"

export const DefaultApiPath = (path:string) => {
  let url = new URL(
    "/api/"+env.PUBLIC_MANIFOLD_DEFAULT_API_VERSION+"/"+path,
    env.PUBLIC_MANIFOLD_UPSTREAM
  )
  console.log(url.href)
  return url.href
}
