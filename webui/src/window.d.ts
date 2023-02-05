import { provider } from 'web3-core';

declare global {
  interface Window {
    ethereum: provider;
  }
}