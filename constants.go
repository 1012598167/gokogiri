package libxml

import("C")

const ( 
        //parser option 
        HTML_PARSE_RECOVER = 1 << 0       //relaxed parsing 
        HTML_PARSE_NOERROR = 1 << 5       //suppress error reports 
        HTML_PARSE_NOWARNING = 1 << 6     //suppress warning reports 
        HTML_PARSE_PEDANTIC = 1 << 7      //pedantic error reporting 
        HTML_PARSE_NOBLANKS = 1 << 8      //remove blank nodes 
        HTML_PARSE_NONET = 1 << 11                //forbid network access 
        HTML_PARSE_COMPACT = 1 << 16      //compact small text nodes 

        //element type 
        XML_ELEMENT_NODE = 1 
        XML_ATTRIBUTE_NODE = 2 
        XML_TEXT_NODE = 3 
        XML_CDATA_SECTION_NODE = 4 
        XML_ENTITY_REF_NODE = 5 
        XML_ENTITY_NODE = 6 
        XML_PI_NODE = 7 
        XML_COMMENT_NODE = 8 
        XML_DOCUMENT_NODE = 9 
        XML_DOCUMENT_TYPE_NODE = 10 
        XML_DOCUMENT_FRAG_NODE = 11 
        XML_NOTATION_NODE = 12 
        XML_HTML_DOCUMENT_NODE = 13 
        XML_DTD_NODE = 14 
        XML_ELEMENT_DECL = 15 
        XML_ATTRIBUTE_DECL = 16 
        XML_ENTITY_DECL = 17 
        XML_NAMESPACE_DECL = 18 
        XML_XINCLUDE_START = 19 
        XML_XINCLUDE_END = 20 
        XML_DOCB_DOCUMENT_NODE = 21 

				XML_SAVE_FORMAT = 1 // format save output
				XML_SAVE_NO_DECL = 2 // drop the xml declaration
				XML_SAVE_NO_EMPTY = 4 // no empty tags
				XML_SAVE_NO_XHTML = 8 // disable XHTML1 specific rules
				XML_SAVE_XHTML = 16 // force XHTML1 specific rules
				XML_SAVE_AS_XML = 32 // force XML serialization on HTML doc
				XML_SAVE_AS_HTML = 64 // force HTML serialization on XML doc
				XML_SAVE_WSNONSIG = 128 // format with non-significant whitespace
)