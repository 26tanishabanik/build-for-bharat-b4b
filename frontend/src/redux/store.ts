import { configureStore } from '@reduxjs/toolkit';
import cartReducer from './cart';
import searchInputActiveReducer from './search';
import errorReducer from './error';

const store = configureStore({
  reducer: {
    cart: cartReducer,
    searchInputActive: searchInputActiveReducer,
    error: errorReducer,
  }
});

export type IRootState = ReturnType<typeof store.getState>;
export default store;