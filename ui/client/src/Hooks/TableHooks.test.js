import { renderHook, act } from '@testing-library/react-hooks';
import React from 'react';
import {
  Router,
} from 'react-router-dom';
import { createBrowserHistory } from 'history';
import { useTableFilterArray, useTableFilterNumber, useTableFilterString } from './TableHooks';

let wrapper;
let history;
beforeEach(() => {
  history = createBrowserHistory();
  // eslint-disable-next-line react/prop-types
  wrapper = ({ children }) => <Router history={history}>{children}</Router>;
});

describe('Hooks: useTableFilterArray tests', () => {
  test('with empty value in url', () => {
    const { result } = renderHook(() => useTableFilterArray('fieldName'), { wrapper });
    expect(result.current[0]).toEqual([]);
    act(() => {
      result.current[1](['1', '2']);
    });
    expect(result.current[0]).toEqual(['1', '2']);
  });

  test('with non empty value in url', () => {
    history.push({
      pathname: '/',
      search: `?${new URLSearchParams({
        fieldName: ['1', '2'],
      })}`,
    });
    const { result } = renderHook(() => useTableFilterArray('fieldName'), { wrapper });
    expect(result.current[0]).toEqual(['1', '2']);
    act(() => {
      result.current[1](['3', '4']);
    });
    expect(result.current[0]).toEqual(['3', '4']);
  });

  test('resetting the value', () => {
    history.push({
      pathname: '/',
      search: `?${new URLSearchParams({
        fieldName: ['1', '2'],
      })}`,
    });
    const { result } = renderHook(() => useTableFilterArray('fieldName'), { wrapper });
    act(() => {
      result.current[1]();
    });
    expect(result.current[0]).not.toBeDefined;
  });
});

describe('Hooks: useTableFilterNumber tests', () => {
  const testNumber = 4;
  test('with empty value in url', () => {
    const { result } = renderHook(() => useTableFilterNumber('fieldName1'), { wrapper });
    expect(result.current[0]).toEqual(null);
    act(() => {
      result.current[1](testNumber);
    });
    expect(result.current[0]).toEqual(testNumber);
  });

  test('with non empty value in url', () => {
    history.push({
      pathname: '/',
      search: `?${new URLSearchParams({
        fieldName1: testNumber,
      })}`,
    });
    const { result } = renderHook(() => useTableFilterNumber('fieldName1'), { wrapper });
    expect(result.current[0]).toEqual(testNumber);
    act(() => {
      result.current[1](testNumber + 1);
    });
    expect(result.current[0]).toEqual(testNumber + 1);
  });

  test('resetting the value', () => {
    history.push({
      pathname: '/',
      search: `?${new URLSearchParams({
        fieldName1: testNumber,
      })}`,
    });
    const { result } = renderHook(() => useTableFilterNumber('fieldName1'), { wrapper });
    act(() => {
      result.current[1]();
    });
    expect(result.current[0]).not.toBeDefined;
  });
});

describe('Hooks: useTableFilterString tests', () => {
  const testString = 'abc';
  test('with empty value in url', () => {
    const { result } = renderHook(() => useTableFilterString('fieldName2'), { wrapper });
    expect(result.current[0]).not.toBeDefined;
    act(() => {
      result.current[1](testString);
    });
    expect(result.current[0]).toEqual(testString);
  });

  test('with non empty value in url', () => {
    history.push({
      pathname: '/',
      search: `?${new URLSearchParams({
        fieldName2: testString,
      })}`,
    });
    const { result } = renderHook(() => useTableFilterString('fieldName2'), { wrapper });
    expect(result.current[0]).toEqual(testString);
    act(() => {
      result.current[1](`${testString}d`);
    });
    expect(result.current[0]).toEqual(`${testString}d`);
  });

  test('resetting the value', () => {
    history.push({
      pathname: '/',
      search: `?${new URLSearchParams({
        fieldName2: testString,
      })}`,
    });
    const { result } = renderHook(() => useTableFilterString('fieldName2'), { wrapper });
    act(() => {
      result.current[1]();
    });
    expect(result.current[0]).not.toBeDefined;
  });
});
