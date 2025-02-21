/**
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import React, { ForwardedRef, forwardRef, PropsWithChildren } from 'react';
import styled from 'styled-components';

import { borderRadius, color, space, SpaceProps } from 'design/system';

export const Icon = forwardRef<HTMLSpanElement, PropsWithChildren<Props>>(
  ({ size = 'medium', children, ...otherProps }, ref) => {
    let iconSize = size;
    if (size === 'small') {
      iconSize = 16;
    }
    if (size === 'medium') {
      iconSize = 20;
    }
    if (size === 'large') {
      iconSize = 24;
    }
    if (size === 'extra-large') {
      iconSize = 32;
    }
    return (
      <StyledIcon {...otherProps} ref={ref}>
        <svg
          fill="currentColor"
          height={iconSize}
          width={iconSize}
          viewBox="0 0 24 24"
        >
          {children}
        </svg>
      </StyledIcon>
    );
  }
);

const StyledIcon = styled.span`
  display: inline-flex;
  align-items: center;
  justify-content: center;

  ${color};
  ${space};
  ${borderRadius};
`;

export type IconSize = 'small' | 'medium' | 'large' | 'extra-large' | number;

/**
 * IconProps are used in each autogenerated icon component.
 */
export type IconProps = SpaceProps & {
  size?: IconSize;
  color?: string;
  title?: string;
  role?: string;
  style?: React.CSSProperties;
  borderRadius?: number;
  onClick?: React.MouseEventHandler;
  disabled?: boolean;
  as?: any;
  to?: string;
  className?: string;
};

/**
 * Props are used on the Icon component, but not in autogenerated icon components.
 */
type Props = IconProps & {
  children?: React.SVGProps<SVGPathElement> | React.SVGProps<SVGPathElement>[];
  a?: any;
  ref?: ForwardedRef<HTMLSpanElement>;
};
