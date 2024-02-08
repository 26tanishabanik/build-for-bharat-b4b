import { Dispatch, UnknownAction } from "@reduxjs/toolkit";
import { setError } from "../redux/error";
import { GetErrorMessage } from "./GetError";

const SearchAutocomplete = async (searchQuery: string, dispatch: Dispatch<UnknownAction>): Promise<string[]> => {
  try {
    const res = await fetch(`${import.meta.env.VITE_BASE_URL}/search/${searchQuery}`);
    const json: { "matchingWords": string[] } = await res.json();
    return json.matchingWords;
  } catch (err) {
    dispatch(setError(GetErrorMessage(err)));
    return [];
  }
}

export default SearchAutocomplete;