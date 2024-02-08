import { Flex, Box, IconButton, Image, Spacer, Badge, Alert, AlertIcon } from '@chakra-ui/react';
import Cart from '../assets/icons/cart.svg';
import Logo from '../assets/images/logo.png';
import { Link } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { IRootState } from '../redux/store';
import { useLocation } from 'react-router-dom';
import Search from './Search';

const Header = () => {
  const location = useLocation();
  const cart = useSelector((state: IRootState) => state.cart);
  const error = useSelector((state: IRootState) => state.error);

  let cartItemCount = 0;
  for (const [_, value] of Object.entries(cart)) {
    cartItemCount += value.Count;
  }

  return (
    <Box w="100%" zIndex="3" position="relative">
      <Box bgColor="white">
        <Flex py={5} mx={3}>
          <Box>
            <Link to="/">
              <Image ml={5} alt='cart' src={Logo} style={{ width: "130px", height: "auto" }} />
            </Link>
          </Box>
          <Spacer />
          {
            location.pathname === "/" ? <></> : <><Search /><Spacer /></>
          }
          <Box mr={5} position="relative">
            <Link to="/cart">
              <IconButton
                p={5}
                aria-label='cart'
                borderRadius={8}
                backgroundColor='transparent'
                icon={<Image boxSize='25px' alt='cart' src={Cart} />}
              />
              <Badge colorScheme='green' position="absolute" left={0.5} bottom={1}>
                {cartItemCount}
              </Badge>
            </Link>
          </Box>
        </Flex>
      </Box>
      {error.error !== "" && <Box mx="10%" mt={5}>
				<Alert status="error" variant="subtle" borderRadius={10}>
					<AlertIcon />
					{error.error}
				</Alert>
			</Box>}
    </Box>
  );
}

export default Header;