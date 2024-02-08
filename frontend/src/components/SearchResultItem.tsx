import { Card, Box, Image, Heading, Stack, CardBody, Text, Button } from '@chakra-ui/react';
import { AddIcon } from '@chakra-ui/icons';
import { useDispatch, useSelector } from 'react-redux';
import { addItemToCart } from '../redux/cart';
import { CartItemAttributes } from '../types/types';
import { IRootState } from '../redux/store';

import { CartButtonWithAddAndRemove } from './CartButtonWithAddAndRemove';

const SearchResultItem = (props: CartItemAttributes) => {
	const dispatch = useDispatch();
	const cart = useSelector((state: IRootState) => state.cart);
	const cartItem = cart[props.productID];

	const AddToCartButton = (props: CartItemAttributes): JSX.Element => {
		return (
			<Button
				leftIcon={<AddIcon />}
				onClick={() => dispatch(addItemToCart({productID: props.productID, cartItemAttributes: props}))}
			>
				Add to Cart
			</Button>
		)
	}

	return (
		<Card
			direction={{ base: 'column', sm: 'row' }}
			variant='elevated'
			height='170px'
			width='100%'
		>
			<Box p={1}>
				<Image
					style={{
						height: "100%",
						width: "100%"
					}}
					borderRadius={5}
					fit='contain'
					src={props.image}
					alt='Caffe Latte'
				/>
			</Box>
			<Stack width='100%'>
				<CardBody position='relative'>
					<Text fontSize="lg" align='right' w="95%" position='absolute'>
						<b>₹ {props.price}</b>
					</Text>
					<Heading size='md'>{props.productName}</Heading>
					<Heading pt={1} size='sm'>{props.seller}</Heading>
					<Text pt={2}>
						{props.quantity}
					</Text>
					<Box display="flex" justifyContent="flex-end">
						{
							cartItem === undefined ? AddToCartButton(props) : CartButtonWithAddAndRemove({CartItemAttributes: props, Count: cartItem?.Count})
						}
					</Box>
				</CardBody>
			</Stack>
		</Card>
	);
};

export default SearchResultItem;