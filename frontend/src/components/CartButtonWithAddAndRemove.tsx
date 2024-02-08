import { ButtonGroup, IconButton, Button } from "@chakra-ui/react";
import { CartItemAttributes } from "../types/types";
import { AddIcon, MinusIcon } from "@chakra-ui/icons";
import { useDispatch } from 'react-redux';
import { addItemToCart, removeItemFromCart } from '../redux/cart';

export const CartButtonWithAddAndRemove = (props: {
  CartItemAttributes: CartItemAttributes,
  Count: number
}): JSX.Element => {
  const dispatch = useDispatch();
  return (
    <ButtonGroup isAttached>
      <IconButton
        aria-label='add'
        icon={<AddIcon />}
        onClick={() => dispatch(addItemToCart({productID: props.CartItemAttributes.productID, cartItemAttributes: props.CartItemAttributes}))} />
      <Button pointerEvents="none">{props.Count}</Button>
      <IconButton
        aria-label='minus'
        icon={<MinusIcon />}
        onClick={() => dispatch(removeItemFromCart({productID: props.CartItemAttributes.productID}))}
      />
    </ButtonGroup>
  );
};
