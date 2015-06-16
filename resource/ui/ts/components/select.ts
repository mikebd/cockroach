// source: components/select.ts
/// <reference path="../typings/mithriljs/mithril.d.ts" />
// Author: Matt Tracy (matt@cockroachlabs.com)
//

/**
 * Components defines reusable components which may be used on multiple pages,
 * or multiple times on the same page.
 */
module Components {
	/**
	 * Select is a basic html option select.
	 */
	export module Select {
		import property = _mithril.MithrilProperty;

		export interface Item {
			value: string;
			text: string;
		}

		export interface Options {
			items: Item[];
			selected: property<string>;
			callback: (string) => void;
		}

		export function controller(options: Options): Options {

			return options;
		}

		export function view(ctrl: Options) {
			var onChangeFn = function(value: string) {
				ctrl.selected(value);
				ctrl.callback(value);
			};
			return m("select", { onchange: m.withAttr("value", onChangeFn) }, [
					ctrl.items.map(function(item) {
						return m('option', { value: item.value }, item.text);
				})
			])
		}
	}
}
