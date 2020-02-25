
(function($){
$.parser.plugins.push("radiobox");
	function init(target){
		var button = $(
				'<span class="radiobox">' +
				'<span class="radiobox-inner">' +
				'<span class="radiobox-checked"></span>' +
				'<span class="radiobox-unchecked"></span>' +
				'<input class="radiobox-value" type="radio">' +
				'</span>' +
				'</span>').insertAfter(target);
		var t = $(target);
		t.addClass('radiobox-f').hide();
		var name = t.attr('name');
		if (name){
			t.removeAttr('name').attr('radioboxName', name);
			button.find('.radiobox-value').attr('name', name);
		}
		return button;
	}
	function createButton(target){
		var state = $.data(target, 'radiobox');
		var opts = state.options;
		var button = state.radiobox;
		var inner = button.find('.radiobox-inner');

		if(opts.label){
			inner.append('<span class="radiobox-label">'+opts.label+'</span>');
		}

		button.find('.radiobox-value')._propAttr('checked', opts.checked);
		button.removeClass('radiobox-disabled').addClass(opts.disabled ? 'radiobox-disabled' : '');
		button.removeClass('radiobox-reversed').addClass(opts.reversed ? 'radiobox-reversed' : '');
		
		checkButton(target, opts.checked);
		setReadonly(target, opts.readonly);
		$(target).radiobox('setValue', opts.value);
	}
	
	function checkButton(target, checked, animate){
		var state = $.data(target, 'radiobox');
		var opts = state.options;
		opts.checked = checked;
		var inner = state.radiobox.find('.radiobox-inner');
		var _checkbox = inner.find('.radiobox-checked');
		var _unchecked = inner.find('.radiobox-unchecked');
		var obj=$('input.radiobox-value[name="'+opts.name+'"]');

		var input = inner.find('.radiobox-value');
		var ck = input.is(':checked');
		if (ck){
			//默认状态
			_checkbox.css("display","block");
			_unchecked.css("display","none");
		}
		if(opts.checked){
			//点击后的执状态
			var s=false;
			obj.each(function(){
				var disabled=$(this).parent('.radiobox-inner').parent('.radiobox.radiobox-disabled').find('.radiobox-value').is(':checked');
				var readonly=$(this).parent('.radiobox-inner').parent('.radiobox.radiobox-readonly').find('.radiobox-value').is(':checked');
				if(disabled||readonly){
					s=true;
					return;
				}
			});
			if(!s){
				obj.parent('.radiobox-inner').find('.radiobox-checked').css("display","none");
				obj.parent('.radiobox-inner').find('.radiobox-unchecked').css("display","block");
				obj.removeAttr('checked');

				$(target).add(input)._propAttr('checked',checked);
				_checkbox.css("display","block");
				_unchecked.css("display","none");
			}

		}



	}
	
	function setDisabled(target, disabled){
		var state = $.data(target, 'radiobox');
		var opts = state.options;
		var button = state.radiobox;
		var input = button.find('.radiobox-value');
		if (disabled){
			opts.disabled = true;
			$(target).add(input).attr('disabled', 'disabled');
			button.addClass('radiobox-disabled');
		} else {
			opts.disabled = false;
			$(target).add(input).removeAttr('disabled');
			button.removeClass('radiobox-disabled');
		}
	}
	
	function setReadonly(target, mode){
		var state = $.data(target, 'radiobox');
		var opts = state.options;
		opts.readonly = mode==undefined ? true : mode;
		state.radiobox.removeClass('radiobox-readonly').addClass(opts.readonly ? 'radiobox-readonly' : '');
	}
	
	function bindEvents(target){
		var state = $.data(target, 'radiobox');
		var opts = state.options;
		state.radiobox.unbind('.radiobox').bind('click.radiobox', function(){
			if (!opts.disabled && !opts.readonly){
				checkButton(target, true, true);
			}
		});
	}
	$.fn.radiobox = function(options, param){
		if (typeof options == 'string'){
			return $.fn.radiobox.methods[options](this, param);
		}
		options = options || {};
		return this.each(function(){
			var state = $.data(this, 'radiobox');
			if (state){
				$.extend(state.options, options);
			} else {
				state = $.data(this, 'radiobox', {
					options: $.extend({}, $.fn.radiobox.defaults, $.fn.radiobox.parseOptions(this), options),
					radiobox: init(this)
				});
			}
			state.options.originalChecked = state.options.checked;
			createButton(this);
			bindEvents(this);
		});
	};
	$.fn.radiobox.methods = {
		options: function(jq){
			var state = jq.data('radiobox');
			return $.extend(state.options, {
				value: state.radiobox.find('.radiobox-value').val()
			});
		},
		enable: function(jq){
			return jq.each(function(){
				setDisabled(this, false);
			});
		},
		disable: function(jq){
			return jq.each(function(){
				setDisabled(this, true);
			});
		},
		readonly: function(jq, mode){
			return jq.each(function(){
				setReadonly(this, mode);
			});
		},
		checked: function(jq, value){
			return jq.each(function(){
				var obj = $(this).data('radiobox');
				var opts = obj.options;
				if (opts.value==value && !opts.disabled && !opts.readonly){
					checkButton(this, true);
				}
			});
		},
		reset: function(jq){
			return jq.each(function(){
				var opts = $(this).radiobox('options');
				checkButton(this, opts.originalChecked);
			});
		},
		setValue: function(jq, value){
			return jq.each(function(){
				$(this).val(value);
				$.data(this, 'radiobox').radiobox.find('.radiobox-value').val(value);
			});
		},
		getValue: function(jq){
			var obj = jq.data('radiobox');
			var opts = obj.options;
			return opts.checked?obj.radiobox.find('.radiobox-value').val():"";
		},
		getText: function(jq){
			var obj = jq.data('radiobox');
			var opts = obj.options;
			return opts.checked?obj.radiobox.find(".radiobox-label").html():"";
		},
		getValues: function(jq){
			var obj = jq.data('radiobox');
			var opts = obj.options;
			return $('input.radiobox-value[name="'+opts.name+'"]:checked').val();
		},
		getTexts: function(jq){
			var obj = jq.data('radiobox');
			var opts = obj.options;
			var obj=$('input.radiobox-value[name="'+opts.name+'"]:checked');
			return obj.parent('.radiobox-inner').find(".radiobox-label").html();
		}
	};
	
	$.fn.radiobox.parseOptions = function(target){
		var t = $(target);
		return $.extend({}, $.parser.parseOptions(target, []), {
			name: (t.attr('name') ? t.attr('name') : undefined),
			label: (t.attr('label') ? t.attr('label') : undefined),
			value: (t.val() || undefined),
			checked: (t.attr('checked') ? true : false),
			disabled: (t.attr('disabled') ? true : false),
			readonly: (t.attr('readonly') ? true : false)
		});
	};
	$.fn.radiobox.defaults = {
		name:"",
		checked: false,
		disabled: false,
		readonly: false,
		label:"",
		value:'',
		onChange: function(checked){}
	};

})(jQuery);
