import { Dispatch, UnknownAction } from "@reduxjs/toolkit";
import { setError } from "../redux/error";
import { GetErrorMessage } from "./GetError";
import { CartItemAttributes } from "../types/types";

const SearchProductGetUUID = async (product: string, dispatch: Dispatch<UnknownAction>): Promise<string> => {
  try {
    const res = await fetch(`${import.meta.env.VITE_BASE_URL}/products/${product}`);
    const json: { "uuid": string } = await res.json();
    return json.uuid;
  } catch (err) {
    dispatch(setError(GetErrorMessage(err)));
    return "";
  }
}

const SearchProductPollResults = async (product: string, uuid: string, dispatch: Dispatch<UnknownAction>): Promise<{"isResult": boolean, "products": CartItemAttributes[]}> => {
  try {
    let url = new URL(`${import.meta.env.VITE_BASE_URL}/products/${product}`);
    url.search = new URLSearchParams({"uuid": uuid}).toString();

    const res = await fetch(url);
    const json: { "isResult": boolean, "products": CartItemAttributes[] } = await res.json();

    return json;
  } catch (err) {
    dispatch(setError(GetErrorMessage(err)));
    return {
      isResult: true,
      products: []
    };
  }
}

export {SearchProductGetUUID, SearchProductPollResults};