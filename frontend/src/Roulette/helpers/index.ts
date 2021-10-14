export const getStaticPath = (path: string, base: string) => {
  // Get the data from the static path
  let formattedPath = path.replace(/^\//, '').replace(/\/$/, '');
  return `${base}/${formattedPath}`;
};
