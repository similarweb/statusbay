// import { renderHook } from '@testing-library/react-hooks';
// import React from 'react';
// import useRealTimeData from './SocketHooks';
// import { SocketIOProvider } from '../context/SocketIOContext';
//
// describe('Hooks: useRealTimeData test', () => {
//   const mockSocket = () => ({
//     on: (event, callback) => {
//       setTimeout(() => {
//         callback({
//           data: {
//             metric1: true,
//           },
//         });
//       });
//     },
//   });
//
//   test('default', async () => {
//     // eslint-disable-next-line react/prop-types
//     const wrapper = ({ children }) => (
//       <SocketIOProvider io={mockSocket}>
//         {children}
//       </SocketIOProvider>
//     );
//     const { result, waitForNextUpdate } = renderHook(() => useRealTimeData('metric1', false), { wrapper });
//     // testing default value(false)
//     expect(result.current).toEqual(false);
//     await waitForNextUpdate();
//     // testing the new value(true)
//     expect(result.current).toEqual(true);
//   });
// });
