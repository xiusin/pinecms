(function($){
$.parser.plugins.push("checkbox");
	function init(target){
		var button = $(
				'<span class="checkbox">' +
				'<span class="checkbox-inner">' +
				'<span class="checkbox-checked"></span>' +
				'<span class="checkbox-unchecked"></span>' +
				'<input class="checkbox-value" type="checkbox">' +
				'</span>' +
				'</span>').insertAfter(target);
		var t = $(target);
		t.addClass('checkbox-f').hide();
		var name = t.attr('name');
		if (name){
			t.removeAttr('name').attr('checkboxName', name);
			button.find('.checkbox-value').attr('name', name);
		}
		return button;
	}
	function createButton(target){
		var state = $.data(target, 'checkbox');
		var opts = state.options;
		var button = state.checkbox;
		var inner = button.find('.checkbox-inner');

		if(opts.label){
			inner.append('<span class="checkbox-label">'+opts.label+'</span>');
		}

		button.find('.checkbox-value')._propAttr('checked', opts.checked);
		button.removeClass('checkbox-disabled').addClass(opts.disabled ? 'checkbox-disabled' : '');
		button.removeClass('checkbox-reversed').addClass(opts.reversed ? 'checkbox-reversed' : '');
		
		checkButton(target, opts.checked);
		setReadonly(target, opts.readonly);
		$(target).checkbox('setValue', opts.value);
	}
	
	function checkButton(target, checked, animate){
		var state = $.data(target, 'checkbox');
		var opts = state.options;
		opts.checked = checked;
		var inner = state.checkbox.find('.checkbox-inner');
		var _checkbox = inner.find('.checkbox-checked');
		var _unchecked = inner.find('.checkbox-unchecked');

		

		var input = inner.find('.checkbox-value');
		var ck = input.is(':checked');
		$(target).add(input)._propAttr('checked', opts.checked);
		if (opts.checked){
			_checkbox.css("display","block");
			_unchecked.css("display","none");
		}else{
			_checkbox.css("display","none");
			_unchecked.css("display","block");
		}
		if (ck != opts.checked){
			opts.onChange.call(target, opts.checked);
		}
	}
	
	function setDisabled(target, disabled){
		var state = $.data(target, 'checkbox');
		var opts = state.options;
		var button = state.checkbox;
		var input = button.find('.checkbox-value');
		if (disabled){
			opts.disabled = true;
			$(target).add(input).attr('disabled', 'disabled');
			button.addClass('checkbox-disabled');
		} else {
			opts.disabled = false;
			$(target).add(input).removeAttr('disabled');
			button.removeClass('checkbox-disabled');
		}
	}
	
	function setReadonly(target, mode){
		var state = $.data(target, 'checkbox');
		var opts = state.options;
		opts.readonly = mode==undefined ? true : mode;
		state.checkbox.removeClass('checkbox-readonly').addClass(opts.readonly ? 'checkbox-readonly' : '');
	}
	
	function bindEvents(target){
		var state = $.data(target, 'checkbox');
		var opts = state.options;
		state.checkbox.unbind('.checkbox').bind('click.checkbox', function(){
			if (!opts.disabled && !opts.readonly){
				checkButton(target, opts.checked ? false : true, true);
			}
		});
	}
	$.fn.checkbox = function(options, param){
		if (typeof options == 'string'){
			return $.fn.checkbox.methods[options](this, param);
		}
		options = options || {};
		return this.each(function(){
			var state = $.data(this, 'checkbox');
			if (state){
				$.extend(state.options, options);
			} else {
				state = $.data(this, 'checkbox', {
					options: $.extend({}, $.fn.checkbox.defaults, $.fn.checkbox.parseOptions(this), options),
					checkbox: init(this)
				});
			}
			state.options.originalChecked = state.options.checked;
			createButton(this);
			bindEvents(this);
		});
	};
	$.fn.checkbox.methods = {
		options: function(jq){
			var state = jq.data('checkbox');
			return $.extend(state.options, {
				value: state.checkbox.find('.checkbox-value').val()
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
		isCheck: function(jq){
			var r=false;
			jq.each(function(){
				var obj = $(this).data('checkbox');
				var opts = obj.options;
				if(opts.checked){
					r=true;
					return;
				}
			});
			return r;
		},
		check: function(jq){
			return jq.each(function(){
				checkButton(this, true);
			});
		},
		uncheck: function(jq){
			return jq.each(function(){
				checkButton(this, false);
			});
		},
		reset: function(jq){
			return jq.each(function(){
				var opts = $(this).checkbox('options');
				checkButton(this, opts.originalChecked);
			});
		},
		setValue: function(jq, value){
			return jq.each(function(){
				$(this).val(value);
				$.data(this, 'checkbox').checkbox.find('.checkbox-value').val(value);
			});
		},
		getValue: function(jq){
			var obj = jq.data('checkbox');
			var opts = obj.options;
			return opts.checked?obj.checkbox.find('.checkbox-value').val():"";
		},
		getText: function(jq){
			var obj = jq.data('checkbox');
			var opts = obj.options;
			return opts.checked?obj.checkbox.find('.checkbox-label').html():"";
		},
		getValues: function(jq){
			var v=[];
			jq.each(function(){
				var obj = $(this).data('checkbox');
				var opts = obj.options;
				if(opts.checked){
					v.push(obj.checkbox.find('.checkbox-value').val()); 
				}
			});
			return v.join(",");
		},
		getTexts:function(jq){
			var v=[];
			jq.each(function(){
				var obj = $(this).data('checkbox');
				var opts = obj.options;
				if(opts.checked){
					v.push(obj.checkbox.find('.checkbox-label').html()); 
				}
			});
			return v.join(",");
		},
		getCount: function(jq){
			var count=0;
			jq.each(function(){
				var obj = $(this).data('checkbox');
				var opts = obj.options;
				if(opts.checked){
					count++;
				}
			});
			return count;
		},
		checkAll:function(jq){
			return jq.each(function(){
				var obj = $(this).data('checkbox');
				var opts = obj.options;
				if (!opts.disabled && !opts.readonly){
					checkButton(this, true);
				}
			});
		},
		uncheckAll:function(jq){
			return jq.each(function(){
				var obj = $(this).data('checkbox');
				var opts = obj.options;
				if (!opts.disabled && !opts.readonly){
					checkButton(this, false);
				}
			});
		},
		reCheck:function(jq){
			return jq.each(function(){
				var obj = $(this).data('checkbox');
				var opts = obj.options;
				if (!opts.disabled && !opts.readonly){
					opts.checked?checkButton(this, false):checkButton(this, true);
				}
			});
		}


	};
	
	$.fn.checkbox.parseOptions = function(target){
		var t = $(target);
		return $.extend({}, $.parser.parseOptions(target, []), {
			label: (t.attr('label') ? t.attr('label') : undefined),
			value: (t.val() || undefined),
			checked: (t.attr('checked') ? true : false),
			disabled: (t.attr('disabled') ? true : false),
			readonly: (t.attr('readonly') ? true : false)
		});
	};
	$.fn.checkbox.defaults = {
		checked: false,
		disabled: false,
		readonly: false,
		label:"",
		value:'1',
		onChange: function(checked){}
	};

})(jQuery);
