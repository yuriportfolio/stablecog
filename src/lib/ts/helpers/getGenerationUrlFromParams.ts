export function getGenerationUrlFromParams({
	prompt,
	negative_prompt,
	model_id,
	scheduler_id,
	width,
	height,
	guidance_scale,
	num_inference_steps,
	seed
}: TGenerationUrlFromParamsParams) {
	const baseUrl = '/?';
	let params: string[] = [];
	if (prompt) params.push(`p=${encodeURIComponent(prompt.text)}`);
	if (negative_prompt) params.push(`np=${encodeURIComponent(negative_prompt.text)}`);
	if (width) params.push(`w=${width}`);
	if (height) params.push(`h=${height}`);
	if (guidance_scale) params.push(`gs=${guidance_scale}`);
	if (num_inference_steps) params.push(`is=${num_inference_steps}`);
	if (seed || seed === 0) params.push(`s=${seed}`);
	if (model_id) params.push(`mi=${model_id}`);
	if (scheduler_id) params.push(`si=${scheduler_id}`);
	return baseUrl + params.join('&');
}

export interface TGenerationUrlFromParamsParams {
	prompt: {
		id: string;
		text: string;
	};
	negative_prompt?: {
		id: string;
		text: string;
	};
	model_id?: string;
	scheduler_id?: string;
	width?: number;
	height?: number;
	seed?: number;
	num_inference_steps?: number;
	guidance_scale?: number;
}
