package rdx 

//import "fmt";

var __RDX_actions = [] int8 { 0, 1, 0, 1, 1, 1, 2, 1, 3, 1, 4, 1, 5, 2, 1, 2, 2, 1, 4, 2, 1, 5, 2, 3, 0, 2, 3, 2, 2, 3, 4, 2, 3, 5, 0 }
var __RDX_key_offsets = [] int16 { 0, 0, 23, 30, 32, 34, 38, 42, 44, 50, 57, 63, 69, 76, 85, 91, 97, 103, 109, 132, 141, 164, 178, 189, 212, 231, 246, 262, 279, 295, 316, 0 }
var __RDX_trans_keys = [] byte { 32, 34, 44, 58, 91, 93, 95, 123, 125, 9, 13, 43, 45, 48, 57, 65, 70, 71, 90, 97, 102, 103, 122, 34, 46, 57, 92, 120, 48, 49, 48, 57, 48, 57, 69, 101, 48, 57, 43, 45, 48, 57, 48, 57, 48, 57, 65, 70, 97, 102, 45, 48, 57, 65, 70, 97, 102, 48, 57, 65, 70, 97, 102, 48, 57, 65, 70, 97, 102, 95, 48, 57, 65, 90, 97, 122, 34, 47, 92, 98, 102, 110, 114, 116, 117, 48, 57, 65, 70, 97, 102, 48, 57, 65, 70, 97, 102, 48, 57, 65, 70, 97, 102, 48, 57, 65, 70, 97, 102, 32, 34, 44, 58, 91, 93, 95, 123, 125, 9, 13, 43, 45, 48, 57, 65, 70, 71, 90, 97, 102, 103, 122, 32, 44, 58, 91, 93, 123, 125, 9, 13, 32, 34, 44, 58, 91, 93, 95, 123, 125, 9, 13, 43, 45, 48, 57, 65, 70, 71, 90, 97, 102, 103, 122, 32, 44, 46, 58, 69, 91, 93, 101, 123, 125, 9, 13, 48, 57, 32, 44, 58, 91, 93, 123, 125, 9, 13, 48, 57, 32, 34, 44, 58, 91, 93, 95, 123, 125, 9, 13, 43, 45, 48, 57, 65, 70, 71, 90, 97, 102, 103, 122, 32, 44, 45, 46, 58, 69, 91, 93, 101, 123, 125, 9, 13, 48, 57, 65, 70, 97, 102, 32, 44, 58, 91, 93, 123, 125, 9, 13, 48, 57, 65, 70, 97, 102, 32, 44, 45, 58, 91, 93, 123, 125, 9, 13, 48, 57, 65, 70, 97, 102, 32, 43, 44, 45, 58, 91, 93, 123, 125, 9, 13, 48, 57, 65, 70, 97, 102, 32, 44, 45, 58, 91, 93, 123, 125, 9, 13, 48, 57, 65, 70, 97, 102, 32, 44, 45, 58, 91, 93, 95, 123, 125, 9, 13, 48, 57, 65, 70, 71, 90, 97, 102, 103, 122, 32, 44, 58, 91, 93, 95, 123, 125, 9, 13, 48, 57, 65, 90, 97, 122, 0 }
var __RDX_single_lengths = [] int8 { 0, 9, 5, 0, 0, 2, 2, 0, 0, 1, 0, 0, 1, 9, 0, 0, 0, 0, 9, 7, 9, 10, 7, 9, 11, 7, 8, 9, 8, 9, 8, 0 }
var __RDX_range_lengths = [] int8 { 0, 7, 1, 1, 1, 1, 1, 1, 3, 3, 3, 3, 3, 0, 3, 3, 3, 3, 7, 1, 7, 2, 2, 7, 4, 4, 4, 4, 4, 6, 4, 0 }
var __RDX_index_offsets = [] int16 { 0, 0, 17, 24, 26, 28, 32, 36, 38, 42, 47, 51, 55, 60, 70, 74, 78, 82, 86, 103, 112, 129, 142, 152, 169, 185, 197, 210, 224, 237, 253, 0 }
var __RDX_cond_targs = [] int8 { 1, 2, 1, 1, 1, 1, 12, 1, 18, 1, 3, 24, 29, 12, 29, 12, 0, 19, 0, 0, 13, 0, 0, 2, 21, 0, 5, 0, 6, 6, 5, 0, 7, 7, 22, 0, 22, 0, 9, 9, 9, 0, 10, 9, 9, 9, 0, 25, 25, 25, 0, 28, 9, 9, 0, 30, 30, 30, 30, 0, 2, 2, 2, 2, 2, 2, 2, 2, 14, 0, 15, 15, 15, 0, 16, 16, 16, 0, 17, 17, 17, 0, 2, 2, 2, 0, 1, 2, 1, 1, 1, 1, 12, 1, 18, 1, 3, 24, 29, 12, 29, 12, 0, 20, 20, 20, 20, 20, 20, 23, 20, 0, 20, 2, 20, 20, 20, 20, 12, 20, 23, 20, 3, 24, 29, 12, 29, 12, 0, 20, 20, 4, 20, 6, 20, 20, 6, 20, 23, 20, 21, 0, 20, 20, 20, 20, 20, 20, 23, 20, 22, 0, 20, 2, 20, 20, 20, 20, 12, 20, 23, 20, 3, 24, 29, 12, 29, 12, 0, 20, 20, 8, 4, 20, 27, 20, 20, 27, 20, 23, 20, 24, 26, 26, 0, 20, 20, 20, 20, 20, 20, 23, 20, 25, 25, 25, 0, 20, 20, 8, 20, 20, 20, 20, 23, 20, 26, 26, 26, 0, 20, 7, 20, 11, 20, 20, 20, 20, 23, 20, 26, 26, 26, 0, 20, 20, 10, 20, 20, 20, 20, 23, 20, 28, 9, 9, 0, 20, 20, 8, 20, 20, 20, 30, 20, 23, 20, 29, 29, 30, 29, 30, 0, 20, 20, 20, 20, 20, 30, 20, 23, 20, 30, 30, 30, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 0 }
var __RDX_cond_actions = [] int8 { 0, 0, 9, 11, 0, 0, 0, 5, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 7, 28, 31, 7, 7, 7, 25, 7, 7, 22, 22, 7, 7, 7, 7, 0, 0, 9, 11, 0, 0, 5, 0, 0, 0, 0, 0, 9, 11, 0, 0, 0, 5, 0, 0, 1, 1, 0, 0, 0, 0, 0, 3, 16, 0, 19, 0, 3, 3, 0, 13, 3, 3, 0, 0, 0, 9, 11, 0, 0, 5, 0, 0, 0, 0, 7, 7, 28, 31, 7, 7, 7, 25, 7, 7, 22, 22, 7, 7, 7, 7, 0, 3, 16, 0, 0, 19, 0, 3, 3, 0, 13, 3, 3, 0, 0, 0, 0, 0, 9, 11, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 9, 0, 11, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 9, 0, 11, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 9, 0, 11, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 9, 0, 11, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 11, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 0, 7, 3, 0, 0, 0, 0, 0, 0, 0 }
var __RDX_eof_trans = [] int16 { 267, 268, 269, 270, 271, 272, 273, 274, 275, 276, 277, 278, 279, 280, 281, 282, 283, 284, 285, 286, 287, 288, 289, 290, 291, 292, 293, 294, 295, 296, 297, 0 }
var _RDX_start int = 1
var _ = _RDX_start
var _RDX_first_final int = 19
var _ = _RDX_first_final
var _RDX_error int = 0
var _ = _RDX_error
var _RDX_en_main int = 1
var _ = _RDX_en_main
func ParseRDX(data []byte) (rdx *RDX, err error) {

	var mark [RdxMaxNesting]int
	nest, cs, p, pe, eof := 0, 0, 0, len(data), len(data)
	
	rdx = &RDX{}
	
	{
		cs=int( _RDX_start );
		
	}
	{
		var _klen int
		var _trans uint  = 0
		var _keys int
		var _acts int
		var _nacts uint
		_resume:
		{
		
		}
		if p == pe && p != eof {
			goto _out;
			
		}
		if p == eof {
			if __RDX_eof_trans[ cs ] > 0 {
				_trans=uint( __RDX_eof_trans[ cs ] ) - 1;
				
			}
			
		} else {
			_keys=int(__RDX_key_offsets[ cs ])
			;
			_trans=uint( __RDX_index_offsets[ cs ] );
			_klen=int( __RDX_single_lengths[ cs ] );
			if _klen > 0 {
				var _lower int = _keys
				var _upper int = _keys + _klen - 1
				var _mid int
				for {
					if _upper < _lower {
						_keys+=_klen;
						_trans+=uint( _klen );
						break;
						
					}
					_mid=_lower + ( ( _upper - _lower ) >> 1 );
					if ( data[ p ] ) < __RDX_trans_keys[ _mid ] {
						_upper=_mid - 1;
						
					} else if ( data[ p ] ) > __RDX_trans_keys[ _mid ] {
						_lower=_mid + 1;
						
					} else {
						_trans+=uint( ( _mid - _keys ) );
						goto _match;
						
					}
					
				}
				
			}
			_klen=int( __RDX_range_lengths[ cs ] );
			if _klen > 0 {
				var _lower int = _keys
				var _upper int = _keys + ( _klen << 1 ) - 2
				var _mid int
				for {
					if _upper < _lower {
						_trans+=uint( _klen );
						break;
						
					}
					_mid=_lower + ( ( ( _upper - _lower ) >> 1 ) & ^ 1 );
					if ( data[ p ] ) < __RDX_trans_keys[ _mid ] {
						_upper=_mid - 2;
						
					} else if ( data[ p ] ) > __RDX_trans_keys[ _mid + 1 ] {
						_lower=_mid + 2;
						
					} else {
						_trans+=uint( ( ( _mid - _keys ) >> 1 ) );
						break;
						
					}
					
				}
				
			}
			_match:
			{
			
			}
			
		}
		cs=int( __RDX_cond_targs[ _trans ] );
		if __RDX_cond_actions[ _trans ] != 0 {
			_acts=int(__RDX_cond_actions[ _trans ])
			;
			_nacts=uint( __RDX_actions[ _acts ] );
			_acts+=1;
			for _nacts > 0 {
				switch __RDX_actions[ _acts ] {
					case 0:
					{mark[nest] = p; }
					
					case 1:
					{rdx.RdxType = RdxInt; 
						rdx.Text = data[mark[nest] : p];
					}
					
					case 2:
					{n := rdx.Nested 
						n = append(n, RDX{Parent: rdx})
						rdx.Nested = n
						rdx.RdxType = RdxMap;
						rdx = &n[len(n)-1]
						nest++; 
					}
					
					case 3:
					{nest--;
						rdx = rdx.Parent;
					}
					
					case 4:
					{n := rdx.Parent.Nested 
						n = append(n, RDX{Parent: rdx.Parent})
						rdx.Parent.Nested = n
						rdx = &n[len(n)-1]
					}
					
					case 5:
					{n := rdx.Parent.Nested 
						n = append(n, RDX{Parent: rdx.Parent})
						rdx.Parent.Nested = n
						rdx = &n[len(n)-1]
					}
					
					
				}
				_nacts-=1;
				_acts+=1;
				
			}
			
		}
		if p == eof {
			if cs >= 19 {
				goto _out;
				
			}
			
		} else {
			if cs != 0 {
				p+=1;
				goto _resume;
				
			}
			
		}
		_out:
		{
		
		}
		
	}
	if cs < _RDX_first_final {
		err = ErrBadRdx
	}
	
	return
}