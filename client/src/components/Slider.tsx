import * as React from 'react';
import Box from '@mui/material/Box';
import Slider from '@mui/material/Slider';

export default function VerticalSlider({ scaleIndex, setValue }: { scaleIndex: number, setValue: any }) {
    function preventHorizontalKeyboardNavigation(event: React.KeyboardEvent) {
        if (event.key === 'ArrowLeft' || event.key === 'ArrowRight') {
            event.preventDefault();
        }
    }
    const handleChange = (event: Event, newValue: number | number[]) => {
        setValue(newValue as number);
    };

    return (
        <Box sx={{ height: 200, position: "absolute", bottom: "7em", right: "2.65em" }}>
            <Slider
                sx={{
                    '& input[type="range"]': {
                        WebkitAppearance: 'slider-vertical',
                    },
                }}
                step={0.1}
                marks
                min={0.1}
                max={10}
                orientation="vertical"
                defaultValue={scaleIndex}
                aria-label="Temperature"
                valueLabelDisplay="auto"
                onChange={handleChange}
                onKeyDown={preventHorizontalKeyboardNavigation}
            />
        </Box>
    );
}
