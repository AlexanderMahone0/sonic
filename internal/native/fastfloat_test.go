/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package native

import (
    `math`
    `strconv`
    `testing`

    `github.com/stretchr/testify/assert`
)

func TestFastFloat_Encode(t *testing.T) {
    var buf [64]byte
    assert.Equal(t, "0"                         , string(buf[:__f64toa(&buf[0], 0)]))
    assert.Equal(t, "0"                         , string(buf[:__f64toa(&buf[0], math.Float64frombits(0x8000000000000000))]))
    assert.Equal(t, "12340000000"               , string(buf[:__f64toa(&buf[0], 1234e7)]))
    assert.Equal(t, "12.34"                     , string(buf[:__f64toa(&buf[0], 1234e-2)]))
    assert.Equal(t, "0.001234"                  , string(buf[:__f64toa(&buf[0], 1234e-6)]))
    assert.Equal(t, "1e30"                      , string(buf[:__f64toa(&buf[0], 1e30)]))
    assert.Equal(t, "1.234e33"                  , string(buf[:__f64toa(&buf[0], 1234e30)]))
    assert.Equal(t, "1.234e308"                 , string(buf[:__f64toa(&buf[0], 1234e305)]))
    assert.Equal(t, "1.234e-317"                , string(buf[:__f64toa(&buf[0], 1234e-320)]))
    assert.Equal(t, "1.7976931348623157e308"    , string(buf[:__f64toa(&buf[0], 1.7976931348623157e308)]))
    assert.Equal(t, "-12340000000"              , string(buf[:__f64toa(&buf[0], -1234e7)]))
    assert.Equal(t, "-12.34"                    , string(buf[:__f64toa(&buf[0], -1234e-2)]))
    assert.Equal(t, "-0.001234"                 , string(buf[:__f64toa(&buf[0], -1234e-6)]))
    assert.Equal(t, "-1e30"                     , string(buf[:__f64toa(&buf[0], -1e30)]))
    assert.Equal(t, "-1.234e33"                 , string(buf[:__f64toa(&buf[0], -1234e30)]))
    assert.Equal(t, "-1.234e308"                , string(buf[:__f64toa(&buf[0], -1234e305)]))
    assert.Equal(t, "-1.234e-317"               , string(buf[:__f64toa(&buf[0], -1234e-320)]))
    assert.Equal(t, "-2.2250738585072014e-308"  , string(buf[:__f64toa(&buf[0], -2.2250738585072014e-308)]))
}

func BenchmarkFastFloat_Encode(b *testing.B) {
    val := -2.2250738585072014e-308
    benchmarks := []struct {
        name string
        test func(*testing.B)
    }{{
        name: "StdLib",
        test: func(b *testing.B) { var buf [64]byte; for i := 0; i < b.N; i++ { strconv.AppendFloat(buf[:], val, 'g', -1, 64) }},
    }, {
        name: "FastFloat",
        test: func(b *testing.B) { var buf [64]byte; for i := 0; i < b.N; i++ { __f64toa(&buf[0], val) }},
    }}
    for _, bm := range benchmarks {
        b.Run(bm.name, bm.test)
    }
}
