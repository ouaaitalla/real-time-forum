export const argument = {
    posts : [],
    limit : 20,
    cursor :0,
    nextcursor :0,
    hasMore : true,
    isLoading : false,
}

export const post = {
    id : 0,
    nickname : "",
    created_at : "",
    title : "",
    content : "",
    likes : 0,
    dislikes : 0,
    comments : 0,
    avatar : ""
}

export let nextCursor = null;
export let hasMore = true;
export let isLoading = false;