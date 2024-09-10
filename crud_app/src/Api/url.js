const domain = "http://0.0.0.0:8000/api/v1"
export const url = {
    login: `${domain}/login`,
    postCreate: `${domain}/posts`,
    postUpdate: (id) => `${domain}/posts/${id}`,
    postDelete: (id) => `${domain}/posts/${id}`,
    postGet: (id) => `${domain}/posts/${id}`,
    postGetAll: `${domain}/posts`,
    profile: `${domain}/profile`,
    postExportCsv: `${domain}/posts/export-csv`
}
