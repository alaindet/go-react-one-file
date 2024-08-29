export function extractQueryParams(urlString: string): Record<string, any> {
  const url = new URL(urlString);
  if (!url.search) {
    return {};
  }

  const rawQueryParams = url.search.slice(1).split('&');
  const queryParams: Record<string, any> = {};
  rawQueryParams.forEach(rawParam => {
    const [key, value] = rawParam.split('=');
    queryParams[key] = value;
  });

  return queryParams;
}
