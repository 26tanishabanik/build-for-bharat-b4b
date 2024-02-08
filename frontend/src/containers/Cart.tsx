import { Container, Flex, Box, Divider, Heading, Center, Spacer, Text } from "@chakra-ui/react";
import { useSelector } from "react-redux";
import { IRootState } from "../redux/store";
import CartItem from "../components/CartItem";

const Cart = () => {
	const cart = useSelector((state: IRootState) => state.cart);

	let subTotal = 0;
	Object.entries(cart).map(mapEntry => {
		const [ _, cartItem ] = mapEntry;
		subTotal += cartItem.CartItemAttributes.price * cartItem.Count;
	})

	const CartView = (): JSX.Element => {
		return (
			<Flex direction='row'>
				<Box w="70%">
					<Flex p={3} mt={3}>
						<Box w="49%">
							<Center>
								<Heading as='h1' fontSize='24px'>Items</Heading>
							</Center>
						</Box>
						<Box w="17%">
							<Center>
								<Heading as='h1' fontSize='24px'>Quantity</Heading>
							</Center>
						</Box>
						<Box w="17%">
							<Center>
								<Heading as='h1' fontSize='24px'>Price</Heading>
							</Center>
						</Box>
						<Box w="17%">
							<Center>
								<Heading as='h1' fontSize='24px'>Remove</Heading>
							</Center>
						</Box>
					</Flex>
					<Divider />
					{
						Object.entries(cart).map(mapEntry => {
							const [ id, cartItem ] = mapEntry;
							return (
								<CartItem
									key={id}
									CartItemAttributes={cartItem.CartItemAttributes}
									Count={cartItem.Count}
								/>
							);
						})
					}
				</Box>
				<Box w="30%" pt={7} pl={7}>
					<Box borderRadius={4} p={5} bgColor='white'>
						<Heading as='h1' fontSize='24px'>Checkout</Heading>
						<Divider py={2} />
						<Flex pt={6}>
							<Heading as='h2' fontSize='18px'>Subtotal:</Heading>
							<Spacer />
							<Text fontSize="17px">₹ {subTotal}</Text>
						</Flex>
						<Flex pt={6}>
							<Heading as='h2' fontSize='18px'>Delivery Charges:</Heading>
							<Spacer />
							<Text fontSize="17px">₹ 0</Text>
						</Flex>
						<Flex pt={6}>
							<Heading as='h2' fontSize='18px'>Taxes:</Heading>
							<Spacer />
							<Text fontSize="17px">₹ 0</Text>
						</Flex>
						<Divider py={2} />
						<Flex py={2}>
							<Spacer />
							<Heading as='h1' fontSize='24px'>₹ {subTotal}</Heading>
						</Flex>
					</Box>
				</Box>
			</Flex>
		);
	};

	const NoItems = (): JSX.Element => {
		return (
			<Center>
				<Heading mt="40vh" as='h1'>
					No items in your cart
				</Heading>
			</Center>
		);
	};

	return (
		<Container maxW='98vw'>
			{
				Object.entries(cart).length !== 0 ? <CartView /> : <NoItems />
			}
		</Container>
	);
}

export default Cart;