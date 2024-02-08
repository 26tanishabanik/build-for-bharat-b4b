import { Flex, Box, Heading, Image, AbsoluteCenter, IconButton, Text } from "@chakra-ui/react";
import { CartItemAttributes } from "../types/types";
import { DeleteIcon } from "@chakra-ui/icons";
import { CartButtonWithAddAndRemove } from "./CartButtonWithAddAndRemove";
import { useDispatch } from "react-redux";
import { removeAllItemsFromCart } from "../redux/cart";

const CartItem = (props: {CartItemAttributes: CartItemAttributes, Count: number}): JSX.Element => {
  const dispatch = useDispatch();
  return (
    <Flex backgroundColor='white' p={3} my={3} borderRadius={5}>
      <Box w="49%">
        <Flex>
          <Image
            alt={props.CartItemAttributes.productName}
            src={props.CartItemAttributes.image}
            style={{ width: "150px", height: "auto" }}
            borderRadius={5}
          />
          <Box ml={4}>
            <Heading size='md'>{props.CartItemAttributes.productName}</Heading>
            <Heading pt={1} size='sm'>{props.CartItemAttributes.seller}</Heading>
            <Text pt={2}>{props.CartItemAttributes.quantity}</Text>
          </Box>
        </Flex>
      </Box>
      <Box w="17%" position='relative'>
        <AbsoluteCenter axis='both'>
          <CartButtonWithAddAndRemove CartItemAttributes={props.CartItemAttributes} Count={props.Count} />
        </AbsoluteCenter>
      </Box>
      <Box w="17%" position='relative'>
        <AbsoluteCenter axis='both'>
          <Heading as='h2' fontSize='18px'>₹ {props.CartItemAttributes.price * props.Count}</Heading>
        </AbsoluteCenter>
      </Box>
      <Box w="17%" position='relative'>
        <AbsoluteCenter axis='both'>
          <IconButton
            aria-label='cart'
            backgroundColor='transparent'
            icon={<DeleteIcon />}
            onClick={() => dispatch(removeAllItemsFromCart({productID: props.CartItemAttributes.productID}))}
          />
        </AbsoluteCenter>
      </Box>
    </Flex>
  );
};

export default CartItem;