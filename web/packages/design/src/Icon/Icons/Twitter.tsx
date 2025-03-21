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

/* MIT License

Copyright (c) 2020 Phosphor Icons

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

import { forwardRef } from 'react';

import { Icon, IconProps } from '../Icon';

/*

THIS FILE IS GENERATED. DO NOT EDIT.

*/

export const Twitter = forwardRef<HTMLSpanElement, IconProps>(
  ({ size = 24, color, ...otherProps }, ref) => (
    <Icon
      size={size}
      color={color}
      className="icon icon-twitter"
      {...otherProps}
      ref={ref}
    >
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M17.7251 5.99895C17.1796 5.5112 16.3437 5.06623 15.106 5.24293C13.7884 5.43098 13.1916 6.37483 12.9147 7.47128C12.7769 8.01683 12.7336 8.55889 12.7265 8.97105C12.7229 9.17528 12.7283 9.34311 12.7343 9.45773C12.7373 9.51493 12.7405 9.55858 12.7428 9.58654L12.7454 9.61636L12.7458 9.62086C12.7682 9.83223 12.6998 10.0436 12.5577 10.2017C12.4155 10.3599 12.2127 10.4502 12 10.4502C11.2118 10.4502 9.96255 10.09 8.49662 9.37416C7.30981 8.79466 5.93292 7.95868 4.46655 6.81768C4.45219 9.73459 5.36154 11.8978 6.32798 13.3777C6.89031 14.2388 7.47491 14.8725 7.9149 15.2879C8.13465 15.4954 8.3175 15.6476 8.44218 15.7458C8.50449 15.7949 8.55216 15.8304 8.58252 15.8524C8.5977 15.8634 8.60854 15.8711 8.61471 15.8754L8.61943 15.8787C8.78462 15.9901 8.89869 16.1627 8.93644 16.3584C8.98837 16.6276 8.87752 16.8497 8.72271 17.0577C8.66352 17.1372 8.57771 17.2486 8.46604 17.384C8.243 17.6544 7.91504 18.0222 7.48797 18.4226C6.93412 18.9418 6.2039 19.5248 5.3104 20.0187C6.42202 20.3528 7.74943 20.4101 9.15542 20.1768C11.2015 19.8371 13.3225 18.894 15.009 17.4335C16.782 15.8982 17.7203 14.0293 18.2155 12.529C18.4627 11.7802 18.5971 11.1298 18.6695 10.6701C18.7057 10.4406 18.7262 10.2596 18.7376 10.1387C18.7433 10.0783 18.7467 10.033 18.7486 10.0044L18.7504 9.97402L18.7506 9.96973C18.758 9.79408 18.8268 9.62605 18.9451 9.49595L20.8046 7.45046H19.2C18.9067 7.45046 18.6404 7.27948 18.5182 7.01289L18.5194 7.01541L18.5143 7.00502C18.5086 6.9937 18.4986 6.97391 18.4839 6.94699C18.4545 6.89299 18.4072 6.81114 18.3411 6.71173C18.2079 6.51146 18.0043 6.24859 17.7251 5.99895ZM19.6354 5.95046C19.621 5.928 19.6059 5.9049 19.5902 5.88123C19.4046 5.60208 19.1207 5.23465 18.7249 4.88075C17.9204 4.16143 16.6563 3.5064 14.894 3.75799C12.7116 4.06946 11.8085 5.72548 11.4603 7.10402C11.2994 7.74137 11.2439 8.35521 11.2295 8.82692C10.7091 8.69341 10.0058 8.44181 9.15478 8.02627C7.78225 7.35608 6.09403 6.29269 4.29062 4.73304C4.07922 4.55021 3.78369 4.5001 3.52384 4.60302C3.26399 4.70594 3.08292 4.94482 3.05403 5.22281C2.63387 9.26677 3.79125 12.2366 5.07206 14.1979C5.70973 15.1744 6.37513 15.8971 6.88514 16.3786C6.97828 16.4665 7.06635 16.5465 7.14825 16.6186C6.97094 16.8216 6.74127 17.0665 6.46206 17.3283C5.71237 18.0311 4.62843 18.8337 3.26285 19.2889C2.99109 19.3795 2.79466 19.617 2.75665 19.9009C2.71864 20.1849 2.84568 20.4656 3.08404 20.6245C4.79305 21.7637 7.10731 22.0372 9.40103 21.6565C11.7071 21.2737 14.0861 20.217 15.991 18.5674C18.0252 16.8058 19.0869 14.6747 19.6399 12.9992C19.9169 12.1601 20.0685 11.4287 20.1513 10.9035C20.1895 10.661 20.2131 10.4617 20.2275 10.3151L23.055 7.20496C23.2547 6.9853 23.306 6.66853 23.1859 6.39704C23.0658 6.12556 22.7969 5.95046 22.5 5.95046H19.6354ZM18.5194 7.01541C18.5196 7.01597 18.5196 7.01601 18.5194 7.01541V7.01541Z"
      />
    </Icon>
  )
);
