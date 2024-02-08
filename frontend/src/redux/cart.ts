import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { CartItemAttributes } from '../types/types';

export interface CartMap {
  [productID: string]: {CartItemAttributes: CartItemAttributes, Count: number}
}

const initialState: CartMap = {};

export const cartSlice = createSlice({
  name: 'cart',
  initialState: initialState,
  reducers: {
    addItemToCart: (state, action: PayloadAction<{productID: string; cartItemAttributes: CartItemAttributes}>) => {
      if (action.payload.productID in state) {
        const cartItem = state[action.payload.productID];
        state[action.payload.productID] = {
          CartItemAttributes: cartItem.CartItemAttributes,
          Count: ++cartItem.Count
        };
      } else {
        state[action.payload.productID] = {
          CartItemAttributes: action.payload.cartItemAttributes,
          Count: 1
        };
      }
    },
    removeItemFromCart: (state, action: PayloadAction<{productID: string}>) => {
      if (action.payload.productID in state) {
        const cartItem = state[action.payload.productID];
        if (cartItem.Count === 1) {
          delete state[action.payload.productID];
        } else {
          state[action.payload.productID] = {
            CartItemAttributes: cartItem.CartItemAttributes,
            Count: --cartItem.Count
          };
        }
      }
    },
    removeAllItemsFromCart: (state, action: PayloadAction<{productID: string}>) => {
      if (action.payload.productID in state) {
        delete state[action.payload.productID];
      }
    }
  }
});

export const { addItemToCart, removeItemFromCart, removeAllItemsFromCart } = cartSlice.actions;

export default cartSlice.reducer;
